package cli

import (
	"github.com/spf13/cobra"
)

// GetCommands returns all available CLI commands
func GetCommands() []*cobra.Command {
	return []*cobra.Command{
		NewCommand(),
		ListCommand(),
		InfoCommand(),
	}
}
