package cli

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/rafa-mori/gocrafter/internal/generator"
	"github.com/rafa-mori/gocrafter/internal/prompt"
	gl "github.com/rafa-mori/gocrafter/logger"
	"github.com/spf13/cobra"
)

// NewCommand creates a new project generation command
func NewCommand() *cobra.Command {
	var (
		template   string
		outputDir  string
		configFile string
		quick      bool
	)

	cmd := &cobra.Command{
		Use:   "new [project-name]",
		Short: "Create a new Go project from template",
		Long: `Create a new Go project using one of the available templates.
This command will guide you through an interactive setup process
or you can use flags for quick generation.`,
		Example: `  # Interactive mode
  gocrafter new

  # Quick mode with template
  gocrafter new my-api --template api-rest

  # Use configuration file
  gocrafter new --config project.json

  # Specify output directory
  gocrafter new my-service --template microservice --output /path/to/projects`,
		RunE: func(cmd *cobra.Command, args []string) error {
			return runNewCommand(args, template, outputDir, configFile, quick)
		},
	}

	cmd.Flags().StringVarP(&template, "template", "t", "", "Template to use (api-rest, cli-tool, microservice, etc.)")
	cmd.Flags().StringVarP(&outputDir, "output", "o", "", "Output directory for the new project")
	cmd.Flags().StringVarP(&configFile, "config", "c", "", "Configuration file to use")
	cmd.Flags().BoolVarP(&quick, "quick", "q", false, "Quick mode with minimal prompts")

	return cmd
}

func runNewCommand(args []string, template, outputDir, configFile string, quick bool) error {
	var config *generator.ProjectConfig
	var err error

	// Load from config file if provided
	if configFile != "" {
		gl.Log("info", fmt.Sprintf("Loading configuration from file: %s", configFile))
		// TODO: Implement config file loading
		return fmt.Errorf("config file loading not yet implemented")
	}

	// Quick mode
	if quick && template != "" {
		gl.Log("info", fmt.Sprintf("Running in quick mode with template: %s", template))
		config, err = prompt.QuickPrompt(template)
		if err != nil {
			return fmt.Errorf("quick prompt failed: %w", err)
		}
	} else if template != "" && len(args) > 0 {
		// Direct mode with template and project name
		config = generator.NewProjectConfig()
		config.Name = args[0]
		config.Template = template
		config.Module = fmt.Sprintf("github.com/user/%s", args[0]) // Default module name
	} else {
		// Interactive mode
		gl.Log("info", "Running interactive mode")
		prompter := prompt.NewInteractivePrompt()
		config, err = prompter.Run()
		if err != nil {
			return fmt.Errorf("interactive prompt failed: %w", err)
		}
	}

	// Set output directory if provided
	if outputDir != "" {
		config.OutputDir = outputDir
	}

	// Validate configuration
	if err := config.Validate(); err != nil {
		return fmt.Errorf("configuration validation failed: %w", err)
	}

	// Get templates path
	templatesPath, err := getTemplatesPath()
	if err != nil {
		return fmt.Errorf("failed to get templates path: %w", err)
	}

	// Create generator and generate project
	gen := generator.NewGenerator(config, templatesPath)
	if err := gen.Generate(); err != nil {
		return fmt.Errorf("project generation failed: %w", err)
	}

	// Success message
	gl.Log("info", fmt.Sprintf("✅ Project generated successfully!"))
	gl.Log("info", fmt.Sprintf("📁 Location: %s", config.GetOutputPath()))
	gl.Log("info", fmt.Sprintf("Next steps:"))
	gl.Log("info", fmt.Sprintf("  cd %s", config.Name))
	gl.Log("info", fmt.Sprintf("  make run    # Start the application"))
	gl.Log("info", fmt.Sprintf("  make test   # Run tests"))
	gl.Log("info", fmt.Sprintf("  make build  # Build the application"))

	return nil
}

func getTemplatesPath() (string, error) {
	// Try to find templates directory relative to the executable
	// First, try relative to current working directory
	if _, err := os.Stat("templates"); err == nil {
		return "templates", nil
	}

	// Try relative to the executable location
	executable, err := os.Executable()
	if err != nil {
		return "", err
	}

	execDir := filepath.Dir(executable)
	templatesPath := filepath.Join(execDir, "templates")
	if _, err := os.Stat(templatesPath); err == nil {
		return templatesPath, nil
	}

	// Try one level up from executable (for development)
	templatesPath = filepath.Join(execDir, "..", "templates")
	if _, err := os.Stat(templatesPath); err == nil {
		return templatesPath, nil
	}

	return "", fmt.Errorf("templates directory not found")
}
