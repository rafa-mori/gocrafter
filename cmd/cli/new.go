package cli

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/rafa-mori/gocrafter/internal/generator"
	"github.com/rafa-mori/gocrafter/internal/prompt"
	"github.com/rafa-mori/gocrafter/internal/types"
	gl "github.com/rafa-mori/gocrafter/logger"
	"github.com/spf13/cobra"
)

// NewCommand creates a new project generation command
func NewCommand() *cobra.Command {
	var (
		template   string
		kit        string
		outputDir  string
		configFile string
		quick      bool
		author     string
		license    string
	)

	cmd := &cobra.Command{
		Use:   "new [project-name]",
		Short: "Create a new Go project from template or kit",
		Long: `Create a new Go project using one of the available templates or kits.
This command will guide you through an interactive setup process
or you can use flags for quick generation.

Templates are built-in project structures, while kits are pluggable
external project templates that can be added from repositories.`,
		Example: `  # Interactive mode
  gocrafter new

  # Quick mode with built-in template
  gocrafter new my-api --template api-rest

  # Use a kit
  gocrafter new my-project --kit golang-web-api

  # Use configuration file
  gocrafter new --config project.json

  # Specify output directory and author
  gocrafter new my-service --kit microservice --output /path/to/projects --author "John Doe"`,
		RunE: func(cmd *cobra.Command, args []string) error {
			return runNewCommand(args, template, kit, outputDir, configFile, quick, author, license)
		},
	}

	cmd.Flags().StringVarP(&template, "template", "t", "", "Built-in template to use (api-rest, cli-tool, microservice, etc.)")
	cmd.Flags().StringVarP(&kit, "kit", "k", "", "Kit to use for project generation")
	cmd.Flags().StringVarP(&outputDir, "output", "o", "", "Output directory for the new project")
	cmd.Flags().StringVarP(&configFile, "config", "c", "", "Configuration file to use")
	cmd.Flags().BoolVarP(&quick, "quick", "q", false, "Quick mode with minimal prompts")
	cmd.Flags().StringVarP(&author, "author", "a", "", "Project author name")
	cmd.Flags().StringVarP(&license, "license", "l", "MIT", "Project license")

	return cmd
}

func runNewCommand(args []string, template, kit, outputDir, configFile string, quick bool, author, license string) error {
	// Validate that both template and kit are not specified
	if template != "" && kit != "" {
		return fmt.Errorf("cannot specify both template and kit. Use either --template or --kit")
	}

	// If kit is specified, use kit generation
	if kit != "" {
		return runKitGeneration(args, kit, outputDir, author, license)
	}

	// Otherwise, use traditional template generation
	return runTemplateGeneration(args, template, outputDir, configFile, quick)
}

func runKitGeneration(args []string, kitName, outputDir, author, license string) error {
	// Validate project name
	if len(args) == 0 {
		return fmt.Errorf("project name is required when using kit generation")
	}
	
	projectName := args[0]
	
	// Initialize kit manager
	kitManager, err := generator.NewKitManager(nil)
	if err != nil {
		return fmt.Errorf("failed to initialize kit manager: %w", err)
	}

	// Create kit generator
	kitGenerator := generator.NewKitGenerator(kitManager)

	// Set output path
	if outputDir == "" {
		outputDir = filepath.Join(".", projectName)
	} else {
		outputDir = filepath.Join(outputDir, projectName)
	}

	// Get kit placeholders
	placeholders, err := kitGenerator.GetKitPlaceholders(kitName)
	if err != nil {
		return fmt.Errorf("failed to get kit placeholders: %w", err)
	}

	// Create placeholder values
	var placeholderValues []types.PlaceholderValue
	
	// Add basic placeholders
	if author != "" {
		placeholderValues = append(placeholderValues, types.PlaceholderValue{
			Name:  "author",
			Value: author,
		})
	}
	
	if license != "" {
		placeholderValues = append(placeholderValues, types.PlaceholderValue{
			Name:  "license",
			Value: license,
		})
	}

	// Prompt for additional placeholders
	prompter := prompt.NewKitPrompt()
	additionalPlaceholders, err := prompter.PromptForPlaceholders(placeholders, placeholderValues)
	if err != nil {
		return fmt.Errorf("failed to prompt for placeholders: %w", err)
	}

	// Merge placeholders
	placeholderValues = append(placeholderValues, additionalPlaceholders...)

	// Create generation request
	req := &types.GenerationRequest{
		KitName:      kitName,
		ProjectName:  projectName,
		OutputPath:   outputDir,
		Placeholders: placeholderValues,
	}

	// Validate request
	if err := kitGenerator.ValidateGenerationRequest(req); err != nil {
		return fmt.Errorf("generation request validation failed: %w", err)
	}

	// Generate project
	if err := kitGenerator.GenerateFromKit(req); err != nil {
		return fmt.Errorf("kit generation failed: %w", err)
	}

	// Success message
	gl.Log("info", "✅ Project generated successfully from kit!")
	gl.Log("info", fmt.Sprintf("📁 Location: %s", outputDir))
	gl.Log("info", fmt.Sprintf("📦 Kit: %s", kitName))
	gl.Log("info", "Next steps:")
	gl.Log("info", fmt.Sprintf("  cd %s", projectName))
	gl.Log("info", "  # Check the generated README.md for specific instructions")

	return nil
}

func runTemplateGeneration(args []string, template, outputDir, configFile string, quick bool) error {
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
	gl.Log("info", "✅ Project generated successfully!")
	gl.Log("info", fmt.Sprintf("📁 Location: %s", config.GetOutputPath()))
	gl.Log("info", "Next steps:")
	gl.Log("info", fmt.Sprintf("  cd %s", config.Name))
	gl.Log("info", "  make run    # Start the application")
	gl.Log("info", "  make test   # Run tests")
	gl.Log("info", "  make build  # Build the application")

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
