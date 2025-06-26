package types

import (
	"time"
)

// Kit represents a pluggable project kit
type Kit struct {
	Name         string            `yaml:"name"`
	Description  string            `yaml:"description"`
	Language     string            `yaml:"language"`
	Version      string            `yaml:"version"`
	Author       string            `yaml:"author"`
	Repository   string            `yaml:"repository"`
	Dependencies []string          `yaml:"dependencies"`
	Placeholders []string          `yaml:"placeholders"`
	Tags         []string          `yaml:"tags"`
	LocalPath    string            `yaml:"-"` // Path where kit is stored locally
	InstallDate  time.Time         `yaml:"-"` // When kit was installed
	Metadata     map[string]string `yaml:"metadata,omitempty"`
}

// KitManager handles kit operations
type KitManager interface {
	// AddKit adds a new kit from repository URL
	AddKit(repoURL string) error
	// RemoveKit removes a kit by name
	RemoveKit(name string) error
	// ListKits returns all available kits
	ListKits() ([]Kit, error)
	// GetKit returns a specific kit by name
	GetKit(name string) (*Kit, error)
	// UpdateKit updates an existing kit
	UpdateKit(name string) error
	// ValidateKit validates kit structure and metadata
	ValidateKit(kitPath string) error
}

// KitConfig represents the configuration for kit management
type KitConfig struct {
	KitsPath    string `yaml:"kits_path"`
	CachePath   string `yaml:"cache_path"`
	AutoUpdate  bool   `yaml:"auto_update"`
	MaxCacheAge int    `yaml:"max_cache_age_days"`
}

// PlaceholderValue represents a placeholder and its value
type PlaceholderValue struct {
	Name        string `yaml:"name"`
	Value       string `yaml:"value"`
	Description string `yaml:"description"`
	Required    bool   `yaml:"required"`
	Default     string `yaml:"default"`
}

// GenerationRequest represents a request to generate a project from a kit
type GenerationRequest struct {
	KitName      string             `yaml:"kit_name"`
	ProjectName  string             `yaml:"project_name"`
	OutputPath   string             `yaml:"output_path"`
	Placeholders []PlaceholderValue `yaml:"placeholders"`
	Options      map[string]string  `yaml:"options"`
}
