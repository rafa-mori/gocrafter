package generator

import (
	"bytes"
	"fmt"
	"io/fs"
	"os"
	"path/filepath"
	"strings"
	"text/template"

	"github.com/rafa-mori/logz"
)

// Generator handles project generation from templates
type Generator struct {
	config        *ProjectConfig
	templateVars  *TemplateVars
	templatesPath string
}

// NewGenerator creates a new project generator
func NewGenerator(config *ProjectConfig, templatesPath string) *Generator {
	return &Generator{
		config:        config,
		templateVars:  config.ToTemplateVars(),
		templatesPath: templatesPath,
	}
}

// Generate creates a new project based on the configuration
func (g *Generator) Generate() error {
	logz.Info("Starting project generation", "project", g.config.Name, "template", g.config.Template)

	// Validate configuration
	if err := g.config.Validate(); err != nil {
		return fmt.Errorf("invalid configuration: %w", err)
	}

	// Create output directory
	outputPath := g.config.GetOutputPath()
	if err := g.createOutputDirectory(outputPath); err != nil {
		return fmt.Errorf("failed to create output directory: %w", err)
	}

	// Get template path
	templatePath := filepath.Join(g.templatesPath, g.config.Template)
	if !g.templateExists(templatePath) {
		return fmt.Errorf("template '%s' not found", g.config.Template)
	}

	// Generate project from template
	if err := g.generateFromTemplate(templatePath, outputPath); err != nil {
		return fmt.Errorf("failed to generate project: %w", err)
	}

	// Post-generation tasks
	if err := g.runPostGeneration(outputPath); err != nil {
		logz.Warn("Post-generation tasks failed", "error", err)
	}

	logz.Info("Project generated successfully", "path", outputPath)
	return nil
}

func (g *Generator) createOutputDirectory(outputPath string) error {
	// Check if directory already exists
	if _, err := os.Stat(outputPath); err == nil {
		return fmt.Errorf("directory '%s' already exists", outputPath)
	}

	// Create directory with all parent directories
	return os.MkdirAll(outputPath, 0755)
}

func (g *Generator) templateExists(templatePath string) bool {
	info, err := os.Stat(templatePath)
	if err != nil {
		return false
	}
	return info.IsDir()
}

func (g *Generator) generateFromTemplate(templatePath, outputPath string) error {
	return filepath.WalkDir(templatePath, func(path string, d fs.DirEntry, err error) error {
		if err != nil {
			return err
		}

		// Get relative path from template root
		relPath, err := filepath.Rel(templatePath, path)
		if err != nil {
			return err
		}

		// Skip template root
		if relPath == "." {
			return nil
		}

		// Process path with template variables
		processedPath := g.processPath(relPath)
		targetPath := filepath.Join(outputPath, processedPath)

		if d.IsDir() {
			// Create directory
			return os.MkdirAll(targetPath, 0755)
		}

		// Process file
		return g.processFile(path, targetPath)
	})
}

func (g *Generator) processPath(path string) string {
	// Replace template variables in path
	processed := strings.ReplaceAll(path, "{{.ProjectName}}", g.templateVars.ProjectName)
	processed = strings.ReplaceAll(processed, "{{.PackageName}}", g.templateVars.PackageName)
	processed = strings.ReplaceAll(processed, "{{.ModuleName}}", g.templateVars.ModuleName)
	return processed
}

func (g *Generator) processFile(sourcePath, targetPath string) error {
	// Read source file
	content, err := os.ReadFile(sourcePath)
	if err != nil {
		return err
	}

	// Check if file should be processed as template
	if g.shouldProcessAsTemplate(sourcePath) {
		processedContent, err := g.processTemplate(string(content))
		if err != nil {
			return fmt.Errorf("failed to process template %s: %w", sourcePath, err)
		}
		content = []byte(processedContent)
	}

	// Write target file
	if err := os.MkdirAll(filepath.Dir(targetPath), 0755); err != nil {
		return err
	}

	return os.WriteFile(targetPath, content, 0644)
}

func (g *Generator) shouldProcessAsTemplate(filePath string) bool {
	// Process certain file types as templates
	ext := strings.ToLower(filepath.Ext(filePath))
	templateExtensions := []string{
		".go", ".mod", ".sum", ".yaml", ".yml", ".json", ".toml",
		".md", ".txt", ".env", ".dockerfile", ".makefile",
	}

	for _, templateExt := range templateExtensions {
		if ext == templateExt {
			return true
		}
	}

	// Also check if filename indicates it's a template
	base := strings.ToLower(filepath.Base(filePath))
	templateFiles := []string{
		"makefile", "dockerfile", "readme", "license", "gitignore",
	}

	for _, templateFile := range templateFiles {
		if strings.Contains(base, templateFile) {
			return true
		}
	}

	return false
}

func (g *Generator) processTemplate(content string) (string, error) {
	// Create template with helper functions
	tmpl := template.New("project").Funcs(template.FuncMap{
		"lower":    strings.ToLower,
		"upper":    strings.ToUpper,
		"title":    strings.Title,
		"contains": strings.Contains,
		"hasFeature": func(feature string) bool {
			for _, f := range g.templateVars.Features {
				if strings.Contains(strings.ToLower(f), strings.ToLower(feature)) {
					return true
				}
			}
			return false
		},
	})

	// Parse template
	tmpl, err := tmpl.Parse(content)
	if err != nil {
		return "", err
	}

	// Execute template
	var buf bytes.Buffer
	err = tmpl.Execute(&buf, g.templateVars)
	if err != nil {
		return "", err
	}

	return buf.String(), nil
}

func (g *Generator) runPostGeneration(outputPath string) error {
	logz.Info("Running post-generation tasks", "path", outputPath)

	// Initialize Go module
	if err := g.initGoModule(outputPath); err != nil {
		return fmt.Errorf("failed to initialize Go module: %w", err)
	}

	// Run go mod tidy
	if err := g.runGoModTidy(outputPath); err != nil {
		return fmt.Errorf("failed to run go mod tidy: %w", err)
	}

	// Format Go code
	if err := g.formatGoCode(outputPath); err != nil {
		return fmt.Errorf("failed to format Go code: %w", err)
	}

	return nil
}

func (g *Generator) initGoModule(outputPath string) error {
	if g.config.Module == "" {
		return nil
	}

	// Check if go.mod already exists
	goModPath := filepath.Join(outputPath, "go.mod")
	if _, err := os.Stat(goModPath); err == nil {
		logz.Info("go.mod already exists, skipping module initialization")
		return nil
	}

	logz.Info("Initializing Go module", "module", g.config.Module)
	return nil // go mod init will be handled by the template
}

func (g *Generator) runGoModTidy(outputPath string) error {
	logz.Info("Running go mod tidy")
	// This would execute: go mod tidy in the output directory
	// For now, we'll leave this as a placeholder since we want to generate templates first
	return nil
}

func (g *Generator) formatGoCode(outputPath string) error {
	logz.Info("Formatting Go code")
	// This would execute: gofmt -w . in the output directory
	// For now, we'll leave this as a placeholder
	return nil
}

// GetTemplateInfo returns information about a template
func GetTemplateInfo(templatesPath, templateName string) (*TemplateInfo, error) {
	templatePath := filepath.Join(templatesPath, templateName)

	if !templateExists(templatePath) {
		return nil, fmt.Errorf("template '%s' not found", templateName)
	}

	// Read template metadata if it exists
	metadataPath := filepath.Join(templatePath, "template.json")
	if _, err := os.Stat(metadataPath); err == nil {
		// Template has metadata file
		return readTemplateMetadata(metadataPath)
	}

	// Default template info
	return &TemplateInfo{
		Name:        templateName,
		Description: fmt.Sprintf("Template for %s projects", templateName),
		Version:     "1.0.0",
	}, nil
}

// TemplateInfo contains information about a template
type TemplateInfo struct {
	Name        string   `json:"name"`
	Description string   `json:"description"`
	Version     string   `json:"version"`
	Author      string   `json:"author,omitempty"`
	Tags        []string `json:"tags,omitempty"`
	Features    []string `json:"features,omitempty"`
}

func templateExists(templatePath string) bool {
	info, err := os.Stat(templatePath)
	if err != nil {
		return false
	}
	return info.IsDir()
}

func readTemplateMetadata(metadataPath string) (*TemplateInfo, error) {
	// This would read and parse the template.json file
	// For now, return a default
	return &TemplateInfo{
		Name:        "unknown",
		Description: "Template",
		Version:     "1.0.0",
	}, nil
}
