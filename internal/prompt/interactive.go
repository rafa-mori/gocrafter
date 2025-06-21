package prompt

import (
	"fmt"
	"path/filepath"
	"strings"

	"github.com/AlecAivazis/survey/v2"
	"github.com/rafa-mori/gocrafter/internal/generator"
	gl "github.com/rafa-mori/gocrafter/logger"
)

// InteractivePrompt handles interactive configuration prompts
type InteractivePrompt struct {
	config *generator.ProjectConfig
}

// NewInteractivePrompt creates a new interactive prompt handler
func NewInteractivePrompt() *InteractivePrompt {
	return &InteractivePrompt{
		config: generator.NewProjectConfig(),
	}
}

// Run executes the interactive prompt flow
func (p *InteractivePrompt) Run() (*generator.ProjectConfig, error) {
	gl.Log("info", "Starting interactive project setup")
	// Print welcome message
	// This is a placeholder for the logger, replace with actual logger if needed
	gl.Log("info", "🚀 Welcome to GoCrafter - Go Project Generator!")
	gl.Log("info", "Let's craft your perfect Go project together...")

	// Project basic info
	if err := p.promptProjectInfo(); err != nil {
		return nil, err
	}

	// Template selection
	if err := p.promptTemplate(); err != nil {
		return nil, err
	}

	// Database selection
	if err := p.promptDatabase(); err != nil {
		return nil, err
	}

	// Additional features
	if err := p.promptFeatures(); err != nil {
		return nil, err
	}

	// DevOps configuration
	if err := p.promptDevOps(); err != nil {
		return nil, err
	}

	// Final confirmation
	if err := p.confirmConfiguration(); err != nil {
		return nil, err
	}

	return p.config, nil
}

func (p *InteractivePrompt) promptProjectInfo() error {
	questions := []*survey.Question{
		{
			Name: "name",
			Prompt: &survey.Input{
				Message: "What's your project name?",
				Help:    "This will be used as the directory name and default package name",
			},
			Validate: survey.Required,
		},
		{
			Name: "module",
			Prompt: &survey.Input{
				Message: "What's your Go module name?",
				Help:    "e.g., github.com/username/project-name",
			},
			Validate: survey.Required,
		},
	}

	answers := struct {
		Name   string
		Module string
	}{}

	err := survey.Ask(questions, &answers)
	if err != nil {
		return err
	}

	p.config.Name = answers.Name
	p.config.Module = answers.Module

	return nil
}

func (p *InteractivePrompt) promptTemplate() error {
	templates := generator.SupportedTemplates()
	templateDescriptions := map[string]string{
		"api-rest":     "REST API with HTTP server (Gin/Fiber)",
		"cli-tool":     "Command-line application with Cobra",
		"microservice": "Microservice with gRPC and HTTP",
		"grpc-service": "Pure gRPC service",
		"worker":       "Background worker/job processor",
		"library":      "Go library/package",
	}

	options := make([]string, len(templates))
	for i, tmpl := range templates {
		desc := templateDescriptions[tmpl]
		options[i] = fmt.Sprintf("%s - %s", tmpl, desc)
	}

	var selected string
	prompt := &survey.Select{
		Message: "What type of project do you want to create?",
		Options: options,
		Help:    "Choose the template that best fits your project needs",
	}

	err := survey.AskOne(prompt, &selected)
	if err != nil {
		return err
	}

	// Extract template name from selection
	p.config.Template = strings.Split(selected, " - ")[0]

	return nil
}

func (p *InteractivePrompt) promptDatabase() error {
	// Only ask for database if it's relevant for the template
	if p.config.Template == "library" || p.config.Template == "cli-tool" {
		return nil
	}

	databases := generator.SupportedDatabases()
	databases = append([]string{"none"}, databases...)

	var selected string
	prompt := &survey.Select{
		Message: "Which database do you want to use?",
		Options: databases,
		Help:    "Select 'none' if you don't need a database",
	}

	err := survey.AskOne(prompt, &selected)
	if err != nil {
		return err
	}

	if selected != "none" {
		p.config.Database = selected

		// Ask for cache if database is selected
		caches := generator.SupportedCaches()
		caches = append([]string{"none"}, caches...)

		var cacheSelected string
		cachePrompt := &survey.Select{
			Message: "Do you want to add a cache layer?",
			Options: caches,
		}

		err = survey.AskOne(cachePrompt, &cacheSelected)
		if err != nil {
			return err
		}

		if cacheSelected != "none" {
			p.config.Cache = cacheSelected
		}
	}

	return nil
}

func (p *InteractivePrompt) promptFeatures() error {
	features := []string{
		"Authentication (JWT)",
		"API Documentation (Swagger)",
		"Health Checks",
		"Metrics (Prometheus)",
		"Distributed Tracing",
		"Rate Limiting",
		"CORS Middleware",
		"Request Validation",
		"Logging Middleware",
		"Recovery Middleware",
	}

	var selected []string
	prompt := &survey.MultiSelect{
		Message: "Which additional features do you want to include?",
		Options: features,
		Help:    "Select all features you want to include in your project",
	}

	err := survey.AskOne(prompt, &selected)
	if err != nil {
		return err
	}

	p.config.Features = selected

	return nil
}

func (p *InteractivePrompt) promptDevOps() error {
	// Docker
	var includeDocker bool
	dockerPrompt := &survey.Confirm{
		Message: "Include Docker configuration?",
		Default: true,
		Help:    "Includes Dockerfile and docker-compose.yml",
	}

	err := survey.AskOne(dockerPrompt, &includeDocker)
	if err != nil {
		return err
	}
	p.config.Docker = includeDocker

	// Kubernetes
	if includeDocker {
		var includeK8s bool
		k8sPrompt := &survey.Confirm{
			Message: "Include Kubernetes manifests?",
			Default: false,
			Help:    "Includes deployment, service, and configmap YAML files",
		}

		err = survey.AskOne(k8sPrompt, &includeK8s)
		if err != nil {
			return err
		}
		p.config.Kubernetes = includeK8s
	}

	// CI/CD
	ciSystems := generator.SupportedCISystems()
	ciSystems = append([]string{"none"}, ciSystems...)

	var selectedCI string
	ciPrompt := &survey.Select{
		Message: "Which CI/CD system do you want to use?",
		Options: ciSystems,
		Help:    "This will generate appropriate workflow files",
	}

	err = survey.AskOne(ciPrompt, &selectedCI)
	if err != nil {
		return err
	}

	if selectedCI != "none" {
		p.config.CI = selectedCI
	}

	return nil
}

func (p *InteractivePrompt) confirmConfiguration() error {
	// Print summary of the configuration
	gl.Log("info", "📋 Project Configuration Summary:")
	gl.Log("info", "  Name: %s\n", p.config.Name)
	gl.Log("info", "  Module: %s\n", p.config.Module)
	gl.Log("info", "  Template: %s\n", p.config.Template)
	if p.config.Database != "" {
		gl.Log("info", fmt.Sprintf("  Database: %s\n", p.config.Database))
	}
	if p.config.Cache != "" {
		gl.Log("info", fmt.Sprintf("  Cache: %s\n", p.config.Cache))
	}
	if len(p.config.Features) > 0 {
		gl.Log("info", fmt.Sprintf("  Features: %s\n", strings.Join(p.config.Features, ", ")))
	}
	gl.Log("info", fmt.Sprintf("  Docker: %t\n", p.config.Docker))
	gl.Log("info", fmt.Sprintf("  Kubernetes: %t\n", p.config.Kubernetes))
	if p.config.CI != "" {
		gl.Log("info", fmt.Sprintf("  CI/CD: %s\n", p.config.CI))
	}

	gl.Log("info", fmt.Sprintf("  Output Directory: %s\n", filepath.Join(".", p.config.Name)))

	var confirm bool
	confirmPrompt := &survey.Confirm{
		Message: "Does this look correct? Proceed with project generation?",
		Default: true,
	}

	err := survey.AskOne(confirmPrompt, &confirm)
	if err != nil {
		return err
	}

	if !confirm {
		return fmt.Errorf("project generation cancelled by user")
	}

	return nil
}

// QuickPrompt runs a simplified prompt for quick project generation
func QuickPrompt(template string) (*generator.ProjectConfig, error) {
	config := generator.NewProjectConfig()
	config.Template = template

	questions := []*survey.Question{
		{
			Name: "name",
			Prompt: &survey.Input{
				Message: "Project name:",
			},
			Validate: survey.Required,
		},
		{
			Name: "module",
			Prompt: &survey.Input{
				Message: "Module name:",
			},
			Validate: survey.Required,
		},
	}

	answers := struct {
		Name   string
		Module string
	}{}

	err := survey.Ask(questions, &answers)
	if err != nil {
		return nil, err
	}

	config.Name = answers.Name
	config.Module = answers.Module

	return config, nil
}
