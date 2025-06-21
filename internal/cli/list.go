package cli

import (
	"fmt"
	"path/filepath"

	"github.com/rafa-mori/gocrafter/internal/generator"
	"github.com/spf13/cobra"
)

// ListCommand creates a command to list available templates
func ListCommand() *cobra.Command {
	cmd := &cobra.Command{
		Use:     "list",
		Aliases: []string{"ls", "templates"},
		Short:   "List available project templates",
		Long:    `List all available project templates with their descriptions.`,
		Example: `  # List all templates
  gocrafter list

  # List templates (alias)
  gocrafter templates`,
		RunE: func(cmd *cobra.Command, args []string) error {
			return runListCommand()
		},
	}

	return cmd
}

func runListCommand() error {
	fmt.Println("🎯 Available Project Templates:")
	fmt.Println()

	templates := generator.SupportedTemplates()
	templateDescriptions := map[string]string{
		"api-rest":     "REST API server with HTTP endpoints (Gin/Fiber)",
		"cli-tool":     "Command-line application with Cobra framework",
		"microservice": "Microservice with gRPC and HTTP endpoints",
		"grpc-service": "Pure gRPC service with protocol buffers",
		"worker":       "Background worker for job processing",
		"library":      "Go library/package for reusable code",
	}

	templateFeatures := map[string][]string{
		"api-rest": {
			"HTTP server with routing",
			"Middleware support",
			"Database integration",
			"Authentication ready",
			"API documentation",
		},
		"cli-tool": {
			"Cobra CLI framework",
			"Subcommands support",
			"Configuration management",
			"Logging system",
			"Cross-platform builds",
		},
		"microservice": {
			"gRPC and HTTP servers",
			"Service discovery",
			"Health checks",
			"Metrics collection",
			"Distributed tracing",
		},
		"grpc-service": {
			"Protocol buffers",
			"gRPC server",
			"Client generation",
			"Streaming support",
			"Service mesh ready",
		},
		"worker": {
			"Job queue integration",
			"Background processing",
			"Retry mechanisms",
			"Monitoring hooks",
			"Graceful shutdown",
		},
		"library": {
			"Package structure",
			"Documentation templates",
			"Testing framework",
			"CI/CD workflows",
			"Version management",
		},
	}

	for _, tmpl := range templates {
		desc := templateDescriptions[tmpl]
		features := templateFeatures[tmpl]

		fmt.Printf("📦 %s\n", tmpl)
		fmt.Printf("   %s\n", desc)

		if len(features) > 0 {
			fmt.Println("   Features:")
			for _, feature := range features {
				fmt.Printf("   • %s\n", feature)
			}
		}

		fmt.Println()
	}

	fmt.Println("💡 Usage:")
	fmt.Println("   gocrafter new --template <template-name>")
	fmt.Println("   gocrafter new  # Interactive mode")

	return nil
}

// InfoCommand creates a command to show detailed information about a template
func InfoCommand() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "info <template-name>",
		Short: "Show detailed information about a template",
		Long:  `Show detailed information about a specific template including its structure and features.`,
		Example: `  # Show info about api-rest template
  gocrafter info api-rest

  # Show info about cli-tool template  
  gocrafter info cli-tool`,
		Args: cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			return runInfoCommand(args[0])
		},
	}

	return cmd
}

func runInfoCommand(templateName string) error {
	// Check if template exists
	supportedTemplates := generator.SupportedTemplates()
	templateExists := false
	for _, tmpl := range supportedTemplates {
		if tmpl == templateName {
			templateExists = true
			break
		}
	}

	if !templateExists {
		return fmt.Errorf("template '%s' not found. Run 'gocrafter list' to see available templates", templateName)
	}

	// Get templates path
	templatesPath, err := getTemplatesPath()
	if err != nil {
		return fmt.Errorf("failed to get templates path: %w", err)
	}

	// Get template info
	info, err := generator.GetTemplateInfo(templatesPath, templateName)
	if err != nil {
		return fmt.Errorf("failed to get template info: %w", err)
	}

	// Display template information
	fmt.Printf("📦 Template: %s\n", info.Name)
	fmt.Printf("📝 Description: %s\n", info.Description)
	fmt.Printf("🏷️  Version: %s\n", info.Version)

	if info.Author != "" {
		fmt.Printf("👤 Author: %s\n", info.Author)
	}

	if len(info.Tags) > 0 {
		fmt.Printf("🏷️  Tags: %s\n", joinStrings(info.Tags, ", "))
	}

	if len(info.Features) > 0 {
		fmt.Println("\n✨ Features:")
		for _, feature := range info.Features {
			fmt.Printf("   • %s\n", feature)
		}
	}

	// Show template structure
	fmt.Println("\n📁 Template Structure:")
	templatePath := filepath.Join(templatesPath, templateName)
	if err := showTemplateStructure(templatePath, ""); err != nil {
		fmt.Printf("   (Unable to show structure: %s)\n", err)
	}

	fmt.Println("\n💡 Usage:")
	fmt.Printf("   gocrafter new --template %s\n", templateName)
	fmt.Printf("   gocrafter new my-project --template %s\n", templateName)

	return nil
}

func showTemplateStructure(templatePath, indent string) error {
	// This would walk through the template directory and show its structure
	// For now, we'll show a placeholder
	fmt.Printf("%s   (Template structure will be shown here)\n", indent)
	return nil
}

func joinStrings(strs []string, separator string) string {
	if len(strs) == 0 {
		return ""
	}

	result := strs[0]
	for i := 1; i < len(strs); i++ {
		result += separator + strs[i]
	}

	return result
}
