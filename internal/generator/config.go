package generator

import (
	"encoding/json"
	"fmt"
	"path/filepath"
	"strings"
)

// ProjectConfig holds the configuration for generating a new project
type ProjectConfig struct {
	Name       string            `json:"name"`
	Module     string            `json:"module"`
	Template   string            `json:"template"`
	Database   string            `json:"database,omitempty"`
	Cache      string            `json:"cache,omitempty"`
	Queue      string            `json:"queue,omitempty"`
	Monitoring []string          `json:"monitoring,omitempty"`
	Docker     bool              `json:"docker"`
	Kubernetes bool              `json:"kubernetes"`
	CI         string            `json:"ci,omitempty"`
	Features   []string          `json:"features,omitempty"`
	Custom     map[string]string `json:"custom,omitempty"`
	OutputDir  string            `json:"output_dir,omitempty"`
}

// TemplateVars contains all variables that will be replaced in templates
type TemplateVars struct {
	ProjectName   string
	ModuleName    string
	PackageName   string
	DatabaseType  string
	CacheType     string
	QueueType     string
	HasDocker     bool
	HasKubernetes bool
	HasMonitoring bool
	CIType        string
	Features      []string
	Custom        map[string]string
}

// NewProjectConfig creates a new project configuration with defaults
func NewProjectConfig() *ProjectConfig {
	return &ProjectConfig{
		Docker:     true,
		Kubernetes: false,
		Features:   []string{},
		Custom:     make(map[string]string),
	}
}

// Validate checks if the configuration is valid
func (c *ProjectConfig) Validate() error {
	if c.Name == "" {
		return fmt.Errorf("project name cannot be empty")
	}

	if c.Module == "" {
		return fmt.Errorf("module name cannot be empty")
	}

	if c.Template == "" {
		return fmt.Errorf("template cannot be empty")
	}

	// Validate project name (should be valid for directory names)
	if strings.ContainsAny(c.Name, "/\\:*?\"<>|") {
		return fmt.Errorf("project name contains invalid characters")
	}

	return nil
}

// ToTemplateVars converts the config to template variables
func (c *ProjectConfig) ToTemplateVars() *TemplateVars {
	packageName := strings.ReplaceAll(c.Name, "-", "")
	packageName = strings.ReplaceAll(packageName, "_", "")
	packageName = strings.ToLower(packageName)

	return &TemplateVars{
		ProjectName:   c.Name,
		ModuleName:    c.Module,
		PackageName:   packageName,
		DatabaseType:  c.Database,
		CacheType:     c.Cache,
		QueueType:     c.Queue,
		HasDocker:     c.Docker,
		HasKubernetes: c.Kubernetes,
		HasMonitoring: len(c.Monitoring) > 0,
		CIType:        c.CI,
		Features:      c.Features,
		Custom:        c.Custom,
	}
}

// GetOutputPath returns the full output path for the project
func (c *ProjectConfig) GetOutputPath() string {
	if c.OutputDir != "" {
		return filepath.Join(c.OutputDir, c.Name)
	}
	return c.Name
}

// ToJSON converts the config to JSON string
func (c *ProjectConfig) ToJSON() (string, error) {
	data, err := json.MarshalIndent(c, "", "  ")
	if err != nil {
		return "", err
	}
	return string(data), nil
}

// FromJSON creates a config from JSON string
func FromJSON(jsonStr string) (*ProjectConfig, error) {
	config := NewProjectConfig()
	err := json.Unmarshal([]byte(jsonStr), config)
	if err != nil {
		return nil, err
	}
	return config, nil
}

// SupportedTemplates returns the list of supported templates
func SupportedTemplates() []string {
	return []string{
		"api-rest",
		"cli-tool",
		"microservice",
		"grpc-service",
		"worker",
		"library",
	}
}

// SupportedDatabases returns the list of supported databases
func SupportedDatabases() []string {
	return []string{
		"postgres",
		"mysql",
		"mongodb",
		"sqlite",
		"redis",
	}
}

// SupportedCaches returns the list of supported cache systems
func SupportedCaches() []string {
	return []string{
		"redis",
		"memcached",
		"in-memory",
	}
}

// SupportedQueues returns the list of supported queue systems
func SupportedQueues() []string {
	return []string{
		"rabbitmq",
		"kafka",
		"redis",
		"nats",
	}
}

// SupportedCISystems returns the list of supported CI systems
func SupportedCISystems() []string {
	return []string{
		"github",
		"gitlab",
		"jenkins",
		"azure",
	}
}
