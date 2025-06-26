package cli

import (
	"fmt"
	"strings"

	"github.com/AlecAivazis/survey/v2"
	"github.com/rafa-mori/gocrafter/internal/generator"
	"github.com/rafa-mori/gocrafter/internal/types"
	gl "github.com/rafa-mori/gocrafter/logger"
	"github.com/spf13/cobra"
)

// KitCommand creates the kit management command
func KitCommand() *cobra.Command {
	cmd := &cobra.Command{
		Use:     "kit",
		Aliases: []string{"k"},
		Short:   "Manage project kits",
		Long:    `Manage pluggable project kits for generating different types of projects.`,
		Example: `  # Add a new kit from repository
  gocrafter kit add https://github.com/user/my-go-kit

  # List all available kits
  gocrafter kit list

  # Remove a kit
  gocrafter kit remove my-go-kit

  # Update a kit
  gocrafter kit update my-go-kit

  # Show kit information
  gocrafter kit info my-go-kit`,
		Annotations: GetDescriptions([]string{"Manage project kits", "Manage pluggable project kits for generating different types of projects."}, false),
	}

	// Add subcommands
	cmd.AddCommand(
		kitAddCommand(),
		kitListCommand(),
		kitRemoveCommand(),
		kitUpdateCommand(),
		kitInfoCommand(),
	)

	return cmd
}

func kitAddCommand() *cobra.Command {
	var force bool

	cmd := &cobra.Command{
		Use:     "add <repository-url>",
		Aliases: []string{"install", "a"},
		Short:   "Add a new kit from repository",
		Long:    `Add a new project kit from a Git repository or archive URL.`,
		Args:    cobra.ExactArgs(1),
		Example: `  # Add kit from GitHub
  gocrafter kit add https://github.com/user/golang-api-kit

  # Add kit from archive
  gocrafter kit add https://example.com/kits/web-kit.tar.gz

  # Force add (overwrite existing)
  gocrafter kit add --force https://github.com/user/my-kit`,
		RunE: func(cmd *cobra.Command, args []string) error {
			return runKitAddCommand(args[0], force)
		},
	}

	cmd.Flags().BoolVarP(&force, "force", "f", false, "Force add kit (overwrite if exists)")
	return cmd
}

func kitListCommand() *cobra.Command {
	var verbose bool

	cmd := &cobra.Command{
		Use:     "list",
		Aliases: []string{"ls", "l"},
		Short:   "List all available kits",
		Long:    `List all installed project kits with their information.`,
		Example: `  # List all kits
  gocrafter kit list

  # List with detailed information
  gocrafter kit list --verbose`,
		RunE: func(cmd *cobra.Command, args []string) error {
			return runKitListCommand(verbose)
		},
	}

	cmd.Flags().BoolVarP(&verbose, "verbose", "v", false, "Show detailed information")
	return cmd
}

func kitRemoveCommand() *cobra.Command {
	var force bool

	cmd := &cobra.Command{
		Use:     "remove <kit-name>",
		Aliases: []string{"rm", "uninstall", "delete"},
		Short:   "Remove a kit",
		Long:    `Remove an installed project kit.`,
		Args:    cobra.ExactArgs(1),
		Example: `  # Remove a kit
  gocrafter kit remove my-go-kit

  # Force remove without confirmation
  gocrafter kit remove --force my-go-kit`,
		RunE: func(cmd *cobra.Command, args []string) error {
			return runKitRemoveCommand(args[0], force)
		},
	}

	cmd.Flags().BoolVarP(&force, "force", "f", false, "Force remove without confirmation")
	return cmd
}

func kitUpdateCommand() *cobra.Command {
	cmd := &cobra.Command{
		Use:     "update <kit-name>",
		Aliases: []string{"upgrade", "u"},
		Short:   "Update a kit",
		Long:    `Update an installed project kit to the latest version.`,
		Args:    cobra.ExactArgs(1),
		Example: `  # Update a specific kit
  gocrafter kit update my-go-kit

  # Update all kits
  gocrafter kit update --all`,
		RunE: func(cmd *cobra.Command, args []string) error {
			return runKitUpdateCommand(args[0])
		},
	}

	return cmd
}

func kitInfoCommand() *cobra.Command {
	cmd := &cobra.Command{
		Use:     "info <kit-name>",
		Aliases: []string{"show", "details"},
		Short:   "Show kit information",
		Long:    `Show detailed information about a specific kit.`,
		Args:    cobra.ExactArgs(1),
		Example: `  # Show kit information
  gocrafter kit info my-go-kit`,
		RunE: func(cmd *cobra.Command, args []string) error {
			return runKitInfoCommand(args[0])
		},
	}

	return cmd
}

// Command implementations

func runKitAddCommand(repoURL string, force bool) error {
	gl.Log("info", fmt.Sprintf("Adding kit from repository: %s", repoURL))

	// Initialize kit manager
	kitManager, err := generator.NewKitManager(nil)
	if err != nil {
		return fmt.Errorf("failed to initialize kit manager: %w", err)
	}

	// Extract kit name for force check
	kitName := extractKitNameFromURL(repoURL)
	if kitName == "" {
		return fmt.Errorf("could not extract kit name from URL")
	}

	// Check if kit already exists
	if !force {
		_, err := kitManager.GetKit(kitName)
		if err == nil {
			// Kit exists, ask for confirmation
			var overwrite bool
			prompt := &survey.Confirm{
				Message: fmt.Sprintf("Kit '%s' already exists. Overwrite?", kitName),
			}
			if err := survey.AskOne(prompt, &overwrite); err != nil {
				return fmt.Errorf("prompt failed: %w", err)
			}
			if !overwrite {
				gl.Log("info", "Kit installation cancelled")
				return nil
			}
			// Remove existing kit
			if err := kitManager.RemoveKit(kitName); err != nil {
				return fmt.Errorf("failed to remove existing kit: %w", err)
			}
		}
	} else {
		// Force mode: remove existing kit if it exists
		if _, err := kitManager.GetKit(kitName); err == nil {
			if err := kitManager.RemoveKit(kitName); err != nil {
				gl.Log("warn", fmt.Sprintf("Failed to remove existing kit: %v", err))
			}
		}
	}

	// Add the kit
	if err := kitManager.AddKit(repoURL); err != nil {
		return fmt.Errorf("failed to add kit: %w", err)
	}

	gl.Log("info", fmt.Sprintf("Kit '%s' added successfully", kitName))
	return nil
}

func runKitListCommand(verbose bool) error {
	// Initialize kit manager
	kitManager, err := generator.NewKitManager(nil)
	if err != nil {
		return fmt.Errorf("failed to initialize kit manager: %w", err)
	}

	// Get all kits
	kits, err := kitManager.ListKits()
	if err != nil {
		return fmt.Errorf("failed to list kits: %w", err)
	}

	if len(kits) == 0 {
		gl.Log("info", "No kits installed")
		gl.Log("info", "Use 'gocrafter kit add <repository-url>' to add a kit")
		return nil
	}

	gl.Log("info", fmt.Sprintf("📦 Installed Kits (%d):", len(kits)))
	gl.Log("info", "")

	for _, kit := range kits {
		if verbose {
			printKitDetailed(kit)
		} else {
			printKitSummary(kit)
		}
	}

	return nil
}

func runKitRemoveCommand(kitName string, force bool) error {
	// Initialize kit manager
	kitManager, err := generator.NewKitManager(nil)
	if err != nil {
		return fmt.Errorf("failed to initialize kit manager: %w", err)
	}

	// Check if kit exists
	_, err = kitManager.GetKit(kitName)
	if err != nil {
		return fmt.Errorf("kit '%s' not found", kitName)
	}

	// Ask for confirmation unless force is used
	if !force {
		var confirm bool
		prompt := &survey.Confirm{
			Message: fmt.Sprintf("Are you sure you want to remove kit '%s'?", kitName),
		}
		if err := survey.AskOne(prompt, &confirm); err != nil {
			return fmt.Errorf("prompt failed: %w", err)
		}
		if !confirm {
			gl.Log("info", "Kit removal cancelled")
			return nil
		}
	}

	// Remove the kit
	if err := kitManager.RemoveKit(kitName); err != nil {
		return fmt.Errorf("failed to remove kit: %w", err)
	}

	return nil
}

func runKitUpdateCommand(kitName string) error {
	// Initialize kit manager
	kitManager, err := generator.NewKitManager(nil)
	if err != nil {
		return fmt.Errorf("failed to initialize kit manager: %w", err)
	}

	// Update the kit
	if err := kitManager.UpdateKit(kitName); err != nil {
		return fmt.Errorf("failed to update kit: %w", err)
	}

	return nil
}

func runKitInfoCommand(kitName string) error {
	// Initialize kit manager
	kitManager, err := generator.NewKitManager(nil)
	if err != nil {
		return fmt.Errorf("failed to initialize kit manager: %w", err)
	}

	// Get kit information
	kit, err := kitManager.GetKit(kitName)
	if err != nil {
		return fmt.Errorf("kit '%s' not found", kitName)
	}

	// Print detailed kit information
	printKitDetailed(*kit)

	// Show placeholders
	kitGenerator := generator.NewKitGenerator(kitManager)
	placeholders, err := kitGenerator.GetKitPlaceholders(kitName)
	if err != nil {
		gl.Log("warn", fmt.Sprintf("Failed to get placeholders: %v", err))
	} else if len(placeholders) > 0 {
		gl.Log("info", "")
		gl.Log("info", "📝 Required Placeholders:")
		for _, placeholder := range placeholders {
			gl.Log("info", fmt.Sprintf("   • %s", placeholder))
		}
	}

	return nil
}

// Helper functions

func extractKitNameFromURL(repoURL string) string {
	if strings.Contains(repoURL, "github.com") {
		parts := strings.Split(repoURL, "/")
		if len(parts) >= 2 {
			name := parts[len(parts)-1]
			return strings.TrimSuffix(name, ".git")
		}
	}
	
	parts := strings.Split(strings.TrimSuffix(repoURL, "/"), "/")
	if len(parts) > 0 {
		return strings.TrimSuffix(parts[len(parts)-1], ".git")
	}
	
	return ""
}

func printKitSummary(kit types.Kit) {
	gl.Log("info", fmt.Sprintf("📦 %s", kit.Name))
	if kit.Description != "" {
		gl.Log("info", fmt.Sprintf("   %s", kit.Description))
	}
	if kit.Language != "" {
		gl.Log("info", fmt.Sprintf("   Language: %s", kit.Language))
	}
	gl.Log("info", "")
}

func printKitDetailed(kit types.Kit) {
	gl.Log("info", fmt.Sprintf("📦 %s", kit.Name))
	
	if kit.Description != "" {
		gl.Log("info", fmt.Sprintf("   Description: %s", kit.Description))
	}
	if kit.Language != "" {
		gl.Log("info", fmt.Sprintf("   Language: %s", kit.Language))
	}
	if kit.Version != "" {
		gl.Log("info", fmt.Sprintf("   Version: %s", kit.Version))
	}
	if kit.Author != "" {
		gl.Log("info", fmt.Sprintf("   Author: %s", kit.Author))
	}
	if kit.Repository != "" {
		gl.Log("info", fmt.Sprintf("   Repository: %s", kit.Repository))
	}
	if kit.LocalPath != "" {
		gl.Log("info", fmt.Sprintf("   Path: %s", kit.LocalPath))
	}
	
	if len(kit.Dependencies) > 0 {
		gl.Log("info", "   Dependencies:")
		for _, dep := range kit.Dependencies {
			gl.Log("info", fmt.Sprintf("     • %s", dep))
		}
	}
	
	if len(kit.Tags) > 0 {
		gl.Log("info", fmt.Sprintf("   Tags: %s", strings.Join(kit.Tags, ", ")))
	}
	
	gl.Log("info", "")
}
