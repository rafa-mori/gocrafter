package main

import (
	"github.com/rafa-mori/gocrafter/internal/cli"
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
	return "craft"
}
func (m *GoCrafter) ShortDescription() string {
	return "GoCrafter is a Go project scaffolding and templating tool."
}
func (m *GoCrafter) LongDescription() string {
	return `GoCrafter: A Go project scaffolding and templating tool.
Create production-ready Go projects with best practices, modern tooling, and customizable templates.`
}
func (m *GoCrafter) Usage() string {
	return "gocrafter [command] [args]"
}
func (m *GoCrafter) Examples() []string {
	return []string{
		"gocrafter new                    # Interactive project creation",
		"gocrafter new --template api-rest # Quick API project",
		"gocrafter list                   # Show available templates",
		"gocrafter info api-rest          # Show template details",
	}
}
func (m *GoCrafter) Active() bool {
	return true
}
func (m *GoCrafter) Module() string {
	return "gocrafter"
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
		Short:   m.ShortDescription(),
		Long:    m.LongDescription(),
	}

	// Add GoCrafter commands
	rtCmd.AddCommand(cli.GetCommands()...)
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
	var printBannerV = os.Getenv("GOCRAFTER_PRINT_BANNER")
	if printBannerV == "" {
		printBannerV = "true"
	}

	return &GoCrafter{
		printBanner: strings.ToLower(printBannerV) == "true",
	}
}
