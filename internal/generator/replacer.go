package generator

import (
	"bytes"
	"fmt"
	"os"
	"path/filepath"
	"regexp"
	"strings"
	"text/template"
	"time"

	"github.com/rafa-mori/gocrafter/internal/types"
	gl "github.com/rafa-mori/gocrafter/logger"
)

// PlaceholderReplacer handles placeholder replacement in templates
type PlaceholderReplacer struct {
	placeholders map[string]string
	funcMap      template.FuncMap
}

// NewPlaceholderReplacer creates a new placeholder replacer
func NewPlaceholderReplacer() *PlaceholderReplacer {
	pr := &PlaceholderReplacer{
		placeholders: make(map[string]string),
		funcMap:      make(template.FuncMap),
	}

	// Add default template functions
	pr.setupDefaultFunctions()
	return pr
}

// SetPlaceholder sets a placeholder value
func (pr *PlaceholderReplacer) SetPlaceholder(name, value string) {
	pr.placeholders[name] = value
}

// SetPlaceholders sets multiple placeholders from a map
func (pr *PlaceholderReplacer) SetPlaceholders(placeholders map[string]string) {
	for name, value := range placeholders {
		pr.placeholders[name] = value
	}
}

// SetPlaceholdersFromRequest sets placeholders from a generation request
func (pr *PlaceholderReplacer) SetPlaceholdersFromRequest(req *types.GenerationRequest) {
	// Set basic placeholders
	pr.SetPlaceholder("project_name", req.ProjectName)
	pr.SetPlaceholder("current_year", fmt.Sprintf("%d", time.Now().Year()))
	
	// Set placeholders from request
	for _, placeholder := range req.Placeholders {
		pr.SetPlaceholder(placeholder.Name, placeholder.Value)
	}
	
	// Set derived placeholders
	pr.setDerivedPlaceholders(req.ProjectName)
}

// ProcessContent processes content with placeholder replacement
func (pr *PlaceholderReplacer) ProcessContent(content string) (string, error) {
	// First pass: simple string replacement for basic placeholders
	processed := pr.simpleReplace(content)
	
	// Second pass: template processing for complex expressions
	return pr.templateProcess(processed)
}

// ProcessPath processes a file path with placeholder replacement
func (pr *PlaceholderReplacer) ProcessPath(path string) string {
	return pr.simpleReplace(path)
}

// GetMissingPlaceholders returns placeholders found in content but not defined
func (pr *PlaceholderReplacer) GetMissingPlaceholders(content string) []string {
	var missing []string
	seen := make(map[string]bool)
	
	// Find all placeholder patterns
	re := regexp.MustCompile(`\{\{([^}]+)\}\}`)
	matches := re.FindAllStringSubmatch(content, -1)
	
	for _, match := range matches {
		if len(match) > 1 {
			placeholder := strings.TrimSpace(match[1])
			
			// Skip template functions and complex expressions
			if strings.Contains(placeholder, " ") || strings.Contains(placeholder, ".") {
				continue
			}
			
			if _, exists := pr.placeholders[placeholder]; !exists && !seen[placeholder] {
				missing = append(missing, placeholder)
				seen[placeholder] = true
			}
		}
	}
	
	return missing
}

// Private methods

func (pr *PlaceholderReplacer) setupDefaultFunctions() {
	pr.funcMap = template.FuncMap{
		// String manipulation
		"upper":    strings.ToUpper,
		"lower":    strings.ToLower,
		"title":    strings.Title,
		"trim":     strings.TrimSpace,
		"replace":  strings.ReplaceAll,
		"contains": strings.Contains,
		"hasPrefix": strings.HasPrefix,
		"hasSuffix": strings.HasSuffix,
		
		// Path manipulation
		"base":     filepath.Base,
		"dir":      filepath.Dir,
		"ext":      filepath.Ext,
		"join":     filepath.Join,
		
		// Conversion
		"kebab": func(s string) string {
			return strings.ToLower(strings.ReplaceAll(s, " ", "-"))
		},
		"snake": func(s string) string {
			return strings.ToLower(strings.ReplaceAll(s, " ", "_"))
		},
		"camel": func(s string) string {
			words := strings.Fields(s)
			if len(words) == 0 {
				return s
			}
			result := strings.ToLower(words[0])
			for _, word := range words[1:] {
				result += strings.Title(strings.ToLower(word))
			}
			return result
		},
		"pascal": func(s string) string {
			words := strings.Fields(s)
			var result strings.Builder
			for _, word := range words {
				result.WriteString(strings.Title(strings.ToLower(word)))
			}
			return result.String()
		},
		
		// Utilities
		"now": func() string {
			return time.Now().Format("2006-01-02 15:04:05")
		},
		"date": func(format string) string {
			return time.Now().Format(format)
		},
		"env": os.Getenv,
		"default": func(defaultValue string, value string) string {
			if value == "" {
				return defaultValue
			}
			return value
		},
	}
}

func (pr *PlaceholderReplacer) setDerivedPlaceholders(projectName string) {
	if projectName == "" {
		return
	}
	
	// Set derived values
	pr.SetPlaceholder("package_name", strings.ToLower(strings.ReplaceAll(projectName, "-", "")))
	pr.SetPlaceholder("module_name", strings.ToLower(strings.ReplaceAll(projectName, " ", "-")))
	pr.SetPlaceholder("class_name", strings.Title(strings.ReplaceAll(projectName, "-", " ")))
	pr.SetPlaceholder("const_name", strings.ToUpper(strings.ReplaceAll(projectName, "-", "_")))
	
	// Go-specific placeholders
	if goVersion := os.Getenv("GO_VERSION"); goVersion != "" {
		pr.SetPlaceholder("go_version", goVersion)
	} else {
		pr.SetPlaceholder("go_version", "1.24") // Default Go version
	}
}

func (pr *PlaceholderReplacer) simpleReplace(content string) string {
	result := content
	
	for name, value := range pr.placeholders {
		// Replace {{placeholder_name}} patterns
		pattern := fmt.Sprintf("{{%s}}", name)
		result = strings.ReplaceAll(result, pattern, value)
		
		// Also support {{.placeholder_name}} patterns for template compatibility
		dotPattern := fmt.Sprintf("{{.%s}}", name)
		result = strings.ReplaceAll(result, dotPattern, value)
	}
	
	return result
}

func (pr *PlaceholderReplacer) templateProcess(content string) (string, error) {
	// Create template with custom functions
	tmpl := template.New("content").Funcs(pr.funcMap)
	
	// Parse template
	tmpl, err := tmpl.Parse(content)
	if err != nil {
		// If template parsing fails, return original content with a warning
		gl.Log("warn", fmt.Sprintf("Template parsing failed, using simple replacement: %v", err))
		return content, nil
	}
	
	// Execute template with placeholders as data
	var buf bytes.Buffer
	if err := tmpl.Execute(&buf, pr.placeholders); err != nil {
		// If template execution fails, return original content with a warning
		gl.Log("warn", fmt.Sprintf("Template execution failed, using simple replacement: %v", err))
		return content, nil
	}
	
	return buf.String(), nil
}

// ExtractPlaceholdersFromKit extracts all placeholders from a kit's templates
func ExtractPlaceholdersFromKit(kitPath string) ([]string, error) {
	var placeholders []string
	seen := make(map[string]bool)
	
	templatesPath := filepath.Join(kitPath, "templates")
	
	err := filepath.Walk(templatesPath, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		
		if info.IsDir() {
			return nil
		}
		
		// Read file content
		content, err := os.ReadFile(path)
		if err != nil {
			gl.Log("warn", fmt.Sprintf("Failed to read file %s: %v", path, err))
			return nil // Continue walking
		}
		
		// Extract placeholders from content
		re := regexp.MustCompile(`\{\{([^}]+)\}\}`)
		matches := re.FindAllStringSubmatch(string(content), -1)
		
		for _, match := range matches {
			if len(match) > 1 {
				placeholder := strings.TrimSpace(match[1])
				
				// Clean up placeholder name (remove dots, spaces, etc.)
				placeholder = strings.TrimPrefix(placeholder, ".")
				if idx := strings.Index(placeholder, " "); idx != -1 {
					placeholder = placeholder[:idx]
				}
				
				if placeholder != "" && !seen[placeholder] {
					placeholders = append(placeholders, placeholder)
					seen[placeholder] = true
				}
			}
		}
		
		return nil
	})
	
	if err != nil {
		return nil, fmt.Errorf("failed to walk templates directory: %w", err)
	}
	
	return placeholders, nil
}
