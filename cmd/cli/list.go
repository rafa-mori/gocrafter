package cli

import (
	"fmt"
	"path/filepath"

	"github.com/rafa-mori/gocrafter/internal/generator"
	gl "github.com/rafa-mori/gocrafter/logger"
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
		Annotations: GetDescriptions([]string{"List available project templates", "List all available project templates with their descriptions."}, false),
		RunE: func(cmd *cobra.Command, args []string) error {
			return runListCommand()
		},
	}

	return cmd
}

func runListCommand() error {
	gl.Log("info", "🎯 Available Project Templates:")

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

		gl.Log("info", fmt.Sprintf("📦 %s\n", tmpl))
		gl.Log("info", fmt.Sprintf("   %s\n", desc))

		if len(features) > 0 {
			gl.Log("info", "   Features:")
			for _, feature := range features {
				gl.Log("info", fmt.Sprintf("   • %s\n", feature))
			}
		}
	}

	gl.Log("info", fmt.Sprintf("💡 Usage:"))
	gl.Log("info", fmt.Sprintf("   gocrafter new --template <template-name>"))
	gl.Log("info", fmt.Sprintf("   gocrafter new  # Interactive mode"))

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
		Args:        cobra.ExactArgs(1),
		Annotations: GetDescriptions([]string{"Show detailed information about a template", "Show detailed information about a specific template including its structure and features."}, false),
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
	gl.Log("info", fmt.Sprintf("📦 Template: %s\n", info.Name))
	gl.Log("info", fmt.Sprintf("📝 Description: %s\n", info.Description))
	gl.Log("info", fmt.Sprintf("🏷️  Version: %s\n", info.Version))

	if info.Author != "" {
		gl.Log("info", fmt.Sprintf("👤 Author: %s\n", info.Author))
	}

	if len(info.Tags) > 0 {
		gl.Log("info", fmt.Sprintf("🏷️  Tags: %s\n", joinStrings(info.Tags, ", ")))
	}

	if len(info.Features) > 0 {
		gl.Log("info", "\n✨ Features:")
		for _, feature := range info.Features {
			gl.Log("info", fmt.Sprintf("   • %s\n", feature))
		}
	}

	// Show template structure
	gl.Log("info", "\n📁 Template Structure:")
	templatePath := filepath.Join(templatesPath, templateName)
	if err := showTemplateStructure(templatePath, ""); err != nil {
		gl.Log("info", fmt.Sprintf("   (Unable to show structure: %s)\n", err))
	}

	gl.Log("info", "\n💡 Usage:")
	gl.Log("info", fmt.Sprintf("   gocrafter new --template %s\n", templateName))
	gl.Log("info", fmt.Sprintf("   gocrafter new my-project --template %s\n", templateName))

	return nil
}

func showTemplateStructure(templatePath, indent string) error {
	// This would walk through the template directory and show its structure
	// For now, we'll show a placeholder
	gl.Log("info", fmt.Sprintf("%s   (Template structure will be shown here)\n", indent))
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
