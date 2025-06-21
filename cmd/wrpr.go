package main

import (
	cc "github.com/rafa-mori/gocrafter/cmd/cli"
	gl "github.com/rafa-mori/gocrafter/logger"
	vs "github.com/rafa-mori/gocrafter/version"
	"github.com/spf13/cobra"

	"os"
	"strings"
)

type GoCrafter struct {
	parentCmdName string
	printBanner   bool
}

func (m *GoCrafter) Alias() string {
	return ""
}
func (m *GoCrafter) ShortDescription() string {
	return "GoCrafter is a minimalistic backend service with Go."
}
func (m *GoCrafter) LongDescription() string {
	return `GoCrafter: A minimalistic backend service with Go.`
}
func (m *GoCrafter) Usage() string {
	return "article [command] [args]"
}
func (m *GoCrafter) Examples() []string {
	return []string{"article some-command",
		"article another-command --option value",
		"article yet-another-command --flag"}
}
func (m *GoCrafter) Active() bool {
	return true
}
func (m *GoCrafter) Module() string {
	return "article"
}
func (m *GoCrafter) Execute() error {
	return m.Command().Execute()
}
func (m *GoCrafter) Command() *cobra.Command {
	gl.Log("debug", "Starting GoCrafter CLI...")

	var rtCmd = &cobra.Command{
		Use:     m.Module(),
		Aliases: []string{m.Alias()},
		Example: m.concatenateExamples(),
		Version: vs.GetVersion(),
		Annotations: cc.GetDescriptions([]string{
			m.LongDescription(),
			m.ShortDescription(),
		}, m.printBanner),
	}

	rtCmd.AddCommand(cc.ServiceCmdList()...)
	rtCmd.AddCommand(vs.CliCommand())

	// Set usage definitions for the command and its subcommands
	setUsageDefinition(rtCmd)
	for _, c := range rtCmd.Commands() {
		setUsageDefinition(c)
		if !strings.Contains(strings.Join(os.Args, " "), c.Use) {
			if c.Short == "" {
				c.Short = c.Annotations["description"]
			}
		}
	}

	return rtCmd
}
func (m *GoCrafter) SetParentCmdName(rtCmd string) {
	m.parentCmdName = rtCmd
}
func (m *GoCrafter) concatenateExamples() string {
	examples := ""
	rtCmd := m.parentCmdName
	if rtCmd != "" {
		rtCmd = rtCmd + " "
	}
	for _, example := range m.Examples() {
		examples += rtCmd + example + "\n  "
	}
	return examples
}
func RegX() *GoCrafter {
	var printBannerV = os.Getenv("GOFORGE_PRINT_BANNER")
	if printBannerV == "" {
		printBannerV = "true"
	}

	return &GoCrafter{
		printBanner: strings.ToLower(printBannerV) == "true",
	}
}
