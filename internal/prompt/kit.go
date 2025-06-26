package prompt

import (
	"fmt"
	"strings"

	"github.com/AlecAivazis/survey/v2"
	"github.com/rafa-mori/gocrafter/internal/types"
	gl "github.com/rafa-mori/gocrafter/logger"
)

// KitPrompt handles prompting for kit-specific placeholders
type KitPrompt struct{}

// NewKitPrompt creates a new kit prompt instance
func NewKitPrompt() *KitPrompt {
	return &KitPrompt{}
}

// PromptForPlaceholders prompts the user for placeholder values
func (kp *KitPrompt) PromptForPlaceholders(required []string, existing []types.PlaceholderValue) ([]types.PlaceholderValue, error) {
	// Create a map of existing values for quick lookup
	existingMap := make(map[string]string)
	for _, pv := range existing {
		existingMap[pv.Name] = pv.Value
	}

	var results []types.PlaceholderValue

	// Filter out placeholders that already have values
	var missingPlaceholders []string
	for _, placeholder := range required {
		if _, exists := existingMap[placeholder]; !exists {
			missingPlaceholders = append(missingPlaceholders, placeholder)
		}
	}

	if len(missingPlaceholders) == 0 {
		return results, nil
	}

	gl.Log("info", "🔧 Configure Kit Placeholders")
	gl.Log("info", "")

	// Prompt for each missing placeholder
	for _, placeholder := range missingPlaceholders {
		value, err := kp.promptForSinglePlaceholder(placeholder)
		if err != nil {
			return nil, fmt.Errorf("failed to prompt for placeholder '%s': %w", placeholder, err)
		}

		if value != "" {
			results = append(results, types.PlaceholderValue{
				Name:  placeholder,
				Value: value,
			})
		}
	}

	return results, nil
}

// promptForSinglePlaceholder prompts for a single placeholder value
func (kp *KitPrompt) promptForSinglePlaceholder(placeholder string) (string, error) {
	// Get default value and prompt message based on placeholder name
	defaultValue, promptMessage := kp.getPlaceholderDefaults(placeholder)

	var value string
	prompt := &survey.Input{
		Message: promptMessage,
		Default: defaultValue,
	}

	if err := survey.AskOne(prompt, &value); err != nil {
		return "", err
	}

	return strings.TrimSpace(value), nil
}

// getPlaceholderDefaults returns default values and prompt messages for common placeholders
func (kp *KitPrompt) getPlaceholderDefaults(placeholder string) (string, string) {
	switch strings.ToLower(placeholder) {
	case "author":
		return "", "Author name:"
	case "license":
		return "MIT", "License (MIT, Apache-2.0, GPL-3.0, etc.):"
	case "description":
		return "", "Project description:"
	case "version":
		return "1.0.0", "Initial version:"
	case "email":
		return "", "Author email:"
	case "repository", "repo":
		return "", "Repository URL:"
	case "homepage", "url":
		return "", "Project homepage:"
	case "keywords":
		return "", "Keywords (comma-separated):"
	case "database", "db":
		return "sqlite", "Database type (sqlite, mysql, postgres, etc.):"
	case "framework":
		return "", "Framework to use:"
	case "port":
		return "8080", "Default port:"
	case "host":
		return "localhost", "Default host:"
	case "namespace":
		return "", "Namespace:"
	case "package":
		return "", "Package name:"
	case "module":
		return "", "Module name:"
	case "go_version":
		return "1.24", "Go version:"
	case "node_version":
		return "18", "Node.js version:"
	case "python_version":
		return "3.9", "Python version:"
	case "java_version":
		return "17", "Java version:"
	case "api_version":
		return "v1", "API version:"
	case "service_name":
		return "", "Service name:"
	case "organization", "org":
		return "", "Organization name:"
	case "team":
		return "", "Team name:"
	case "environment", "env":
		return "development", "Environment (development, staging, production):"
	case "region":
		return "us-east-1", "AWS region:"
	case "cluster":
		return "", "Cluster name:"
	case "domain":
		return "", "Domain name:"
	case "subdomain":
		return "", "Subdomain:"
	case "protocol":
		return "http", "Protocol (http, https):"
	case "container_registry":
		return "docker.io", "Container registry:"
	case "image_name":
		return "", "Docker image name:"
	case "dockerfile":
		return "Dockerfile", "Dockerfile name:"
	case "makefile":
		return "Makefile", "Makefile name:"
	case "ci_provider":
		return "github", "CI provider (github, gitlab, jenkins, etc.):"
	case "monitoring":
		return "prometheus", "Monitoring system (prometheus, datadog, etc.):"
	case "logging":
		return "logrus", "Logging library:"
	case "testing_framework":
		return "testify", "Testing framework:"
	case "orm":
		return "gorm", "ORM library:"
	case "router":
		return "gin", "HTTP router (gin, echo, mux, etc.):"
	case "cache":
		return "redis", "Cache system (redis, memcached, etc.):"
	case "queue":
		return "redis", "Queue system (redis, rabbitmq, kafka, etc.):"
	case "storage":
		return "local", "Storage type (local, s3, gcs, etc.):"
	default:
		// For unknown placeholders, create a generic prompt
		promptMessage := fmt.Sprintf("%s:", kp.humanizePlaceholderName(placeholder))
		return "", promptMessage
	}
}

// humanizePlaceholderName converts placeholder names to human-readable format
func (kp *KitPrompt) humanizePlaceholderName(placeholder string) string {
	// Replace underscores and hyphens with spaces
	humanized := strings.ReplaceAll(placeholder, "_", " ")
	humanized = strings.ReplaceAll(humanized, "-", " ")
	
	// Capitalize first letter of each word
	words := strings.Fields(humanized)
	for i, word := range words {
		if len(word) > 0 {
			words[i] = strings.ToUpper(word[:1]) + strings.ToLower(word[1:])
		}
	}
	
	return strings.Join(words, " ")
}

// PromptForKitSelection prompts user to select from available kits
func (kp *KitPrompt) PromptForKitSelection(kits []types.Kit) (*types.Kit, error) {
	if len(kits) == 0 {
		return nil, fmt.Errorf("no kits available")
	}

	// Create options for selection
	var options []string
	kitMap := make(map[string]*types.Kit)

	for i := range kits {
		option := fmt.Sprintf("%s - %s", kits[i].Name, kits[i].Description)
		if kits[i].Language != "" {
			option += fmt.Sprintf(" (%s)", kits[i].Language)
		}
		options = append(options, option)
		kitMap[option] = &kits[i]
	}

	var selected string
	prompt := &survey.Select{
		Message: "Select a kit:",
		Options: options,
	}

	if err := survey.AskOne(prompt, &selected); err != nil {
		return nil, err
	}

	return kitMap[selected], nil
}

// ConfirmGeneration asks for final confirmation before generating
func (kp *KitPrompt) ConfirmGeneration(kitName, projectName, outputPath string, placeholders []types.PlaceholderValue) (bool, error) {
	gl.Log("info", "")
	gl.Log("info", "📋 Generation Summary:")
	gl.Log("info", fmt.Sprintf("   Kit: %s", kitName))
	gl.Log("info", fmt.Sprintf("   Project: %s", projectName))
	gl.Log("info", fmt.Sprintf("   Output: %s", outputPath))
	
	if len(placeholders) > 0 {
		gl.Log("info", "   Placeholders:")
		for _, p := range placeholders {
			if p.Value != "" {
				gl.Log("info", fmt.Sprintf("     %s: %s", p.Name, p.Value))
			}
		}
	}
	
	gl.Log("info", "")

	var confirm bool
	prompt := &survey.Confirm{
		Message: "Generate project with these settings?",
		Default: true,
	}

	if err := survey.AskOne(prompt, &confirm); err != nil {
		return false, err
	}

	return confirm, nil
}
