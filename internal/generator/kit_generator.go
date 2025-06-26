package generator

import (
	"fmt"
	"io/fs"
	"os"
	"os/exec"
	"path/filepath"
	"strings"

	"github.com/rafa-mori/gocrafter/internal/types"
	gl "github.com/rafa-mori/gocrafter/logger"
)

// KitGenerator generates projects from kits
type KitGenerator struct {
	kitManager *KitManagerImpl
	replacer   *PlaceholderReplacer
}

// NewKitGenerator creates a new kit-based project generator
func NewKitGenerator(kitManager *KitManagerImpl) *KitGenerator {
	return &KitGenerator{
		kitManager: kitManager,
		replacer:   NewPlaceholderReplacer(),
	}
}

// GenerateFromKit generates a project from a kit
func (kg *KitGenerator) GenerateFromKit(req *types.GenerationRequest) error {
	gl.Log("info", fmt.Sprintf("Generating project '%s' from kit '%s'", req.ProjectName, req.KitName))

	// Get kit
	kit, err := kg.kitManager.GetKit(req.KitName)
	if err != nil {
		return fmt.Errorf("failed to get kit: %w", err)
	}

	// Validate output path
	if err := kg.validateOutputPath(req.OutputPath); err != nil {
		return fmt.Errorf("invalid output path: %w", err)
	}

	// Setup placeholders
	kg.replacer.SetPlaceholdersFromRequest(req)
	
	// Add kit-specific placeholders
	if err := kg.setupKitPlaceholders(kit, req); err != nil {
		return fmt.Errorf("failed to setup kit placeholders: %w", err)
	}

	// Generate project structure
	templatesPath := filepath.Join(kit.LocalPath, "templates")
	if err := kg.generateFromTemplates(templatesPath, req.OutputPath); err != nil {
		return fmt.Errorf("failed to generate from templates: %w", err)
	}

	// Run post-generation script if exists
	if err := kg.runPostGenerationScript(kit.LocalPath, req.OutputPath); err != nil {
		gl.Log("warn", fmt.Sprintf("Post-generation script failed: %v", err))
	}

	gl.Log("info", fmt.Sprintf("Project '%s' generated successfully at: %s", req.ProjectName, req.OutputPath))
	return nil
}

// ValidateGenerationRequest validates a generation request
func (kg *KitGenerator) ValidateGenerationRequest(req *types.GenerationRequest) error {
	// Validate required fields
	if req.KitName == "" {
		return fmt.Errorf("kit name is required")
	}
	if req.ProjectName == "" {
		return fmt.Errorf("project name is required")
	}
	if req.OutputPath == "" {
		return fmt.Errorf("output path is required")
	}

	// Check if kit exists
	_, err := kg.kitManager.GetKit(req.KitName)
	if err != nil {
		return fmt.Errorf("kit validation failed: %w", err)
	}

	return nil
}

// GetKitPlaceholders returns all placeholders required by a kit
func (kg *KitGenerator) GetKitPlaceholders(kitName string) ([]string, error) {
	kit, err := kg.kitManager.GetKit(kitName)
	if err != nil {
		return nil, fmt.Errorf("failed to get kit: %w", err)
	}

	// Get placeholders from kit metadata
	placeholders := make([]string, len(kit.Placeholders))
	copy(placeholders, kit.Placeholders)

	// Extract additional placeholders from templates
	templatePlaceholders, err := ExtractPlaceholdersFromKit(kit.LocalPath)
	if err != nil {
		gl.Log("warn", fmt.Sprintf("Failed to extract placeholders from templates: %v", err))
	} else {
		// Merge with metadata placeholders
		seen := make(map[string]bool)
		for _, p := range placeholders {
			seen[p] = true
		}
		
		for _, p := range templatePlaceholders {
			if !seen[p] {
				placeholders = append(placeholders, p)
				seen[p] = true
			}
		}
	}

	return placeholders, nil
}

// Private methods

func (kg *KitGenerator) validateOutputPath(outputPath string) error {
	// Check if path already exists
	if _, err := os.Stat(outputPath); err == nil {
		return fmt.Errorf("output path '%s' already exists", outputPath)
	}

	// Check if parent directory exists and is writable
	parentDir := filepath.Dir(outputPath)
	if info, err := os.Stat(parentDir); err != nil {
		if os.IsNotExist(err) {
			// Try to create parent directories
			if err := os.MkdirAll(parentDir, 0755); err != nil {
				return fmt.Errorf("failed to create parent directories: %w", err)
			}
		} else {
			return fmt.Errorf("failed to access parent directory: %w", err)
		}
	} else if !info.IsDir() {
		return fmt.Errorf("parent path is not a directory: %s", parentDir)
	}

	return nil
}

func (kg *KitGenerator) setupKitPlaceholders(kit *types.Kit, req *types.GenerationRequest) error {
	// Set kit metadata as placeholders
	kg.replacer.SetPlaceholder("kit_name", kit.Name)
	kg.replacer.SetPlaceholder("kit_version", kit.Version)
	kg.replacer.SetPlaceholder("kit_author", kit.Author)
	
	// Set default values for common placeholders if not provided
	kg.setDefaultPlaceholders(req)

	return nil
}

func (kg *KitGenerator) setDefaultPlaceholders(req *types.GenerationRequest) {
	defaults := map[string]string{
		"author":      "Developer",
		"license":     "MIT",
		"description": fmt.Sprintf("A project generated from kit %s", req.KitName),
		"version":     "1.0.0",
	}

	// Only set defaults if not already provided
	provided := make(map[string]bool)
	for _, p := range req.Placeholders {
		provided[p.Name] = true
	}

	for name, value := range defaults {
		if !provided[name] {
			kg.replacer.SetPlaceholder(name, value)
		}
	}
}

func (kg *KitGenerator) generateFromTemplates(templatesPath, outputPath string) error {
	return filepath.WalkDir(templatesPath, func(path string, d fs.DirEntry, err error) error {
		if err != nil {
			return err
		}

		// Get relative path from templates root
		relPath, err := filepath.Rel(templatesPath, path)
		if err != nil {
			return err
		}

		// Skip templates root
		if relPath == "." {
			return nil
		}

		// Process path with placeholders
		processedPath := kg.replacer.ProcessPath(relPath)
		targetPath := filepath.Join(outputPath, processedPath)

		if d.IsDir() {
			// Create directory
			return os.MkdirAll(targetPath, 0755)
		}

		// Process file
		return kg.processTemplateFile(path, targetPath)
	})
}

func (kg *KitGenerator) processTemplateFile(sourcePath, targetPath string) error {
	// Read source file
	content, err := os.ReadFile(sourcePath)
	if err != nil {
		return fmt.Errorf("failed to read source file: %w", err)
	}

	// Process content with placeholders if it's a template file
	processedContent := string(content)
	if kg.shouldProcessAsTemplate(sourcePath) {
		processedContent, err = kg.replacer.ProcessContent(string(content))
		if err != nil {
			return fmt.Errorf("failed to process template content: %w", err)
		}
	}

	// Ensure target directory exists
	if err := os.MkdirAll(filepath.Dir(targetPath), 0755); err != nil {
		return fmt.Errorf("failed to create target directory: %w", err)
	}

	// Write processed content to target file
	if err := os.WriteFile(targetPath, []byte(processedContent), 0644); err != nil {
		return fmt.Errorf("failed to write target file: %w", err)
	}

	gl.Log("debug", fmt.Sprintf("Generated file: %s", targetPath))
	return nil
}

func (kg *KitGenerator) shouldProcessAsTemplate(filePath string) bool {
	// Check file extension
	ext := strings.ToLower(filepath.Ext(filePath))
	templateExtensions := []string{
		".go", ".mod", ".sum", ".yaml", ".yml", ".json", ".toml",
		".md", ".txt", ".env", ".dockerfile", ".makefile", ".sh",
		".js", ".ts", ".jsx", ".tsx", ".css", ".scss", ".html",
		".py", ".rs", ".java", ".kt", ".cpp", ".c", ".h",
	}

	for _, templateExt := range templateExtensions {
		if ext == templateExt {
			return true
		}
	}

	// Check filename patterns
	base := strings.ToLower(filepath.Base(filePath))
	templateFiles := []string{
		"makefile", "dockerfile", "readme", "license", "gitignore",
		"changelog", "contributing", "notice", "authors",
	}

	for _, templateFile := range templateFiles {
		if strings.Contains(base, templateFile) {
			return true
		}
	}

	// Check for .tpl extension
	if strings.HasSuffix(filePath, ".tpl") {
		return true
	}

	return false
}

func (kg *KitGenerator) runPostGenerationScript(kitPath, outputPath string) error {
	scriptPath := filepath.Join(kitPath, "scaffold.sh")
	
	// Check if script exists
	if _, err := os.Stat(scriptPath); os.IsNotExist(err) {
		return nil // No script to run
	}

	gl.Log("info", "Running post-generation script...")

	// Make script executable
	if err := os.Chmod(scriptPath, 0755); err != nil {
		return fmt.Errorf("failed to make script executable: %w", err)
	}

	// Run script in the output directory
	cmd := exec.Command("/bin/bash", scriptPath)
	cmd.Dir = outputPath
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	// Set environment variables for the script
	cmd.Env = append(os.Environ(),
		fmt.Sprintf("GOCRAFTER_PROJECT_PATH=%s", outputPath),
		fmt.Sprintf("GOCRAFTER_KIT_PATH=%s", kitPath),
	)

	if err := cmd.Run(); err != nil {
		return fmt.Errorf("post-generation script failed: %w", err)
	}

	gl.Log("info", "Post-generation script completed successfully")
	return nil
}
