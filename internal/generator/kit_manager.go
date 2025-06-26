package generator

import (
	"archive/tar"
	"compress/gzip"
	"fmt"
	"io"
	"net/http"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
	"time"

	"github.com/rafa-mori/gocrafter/internal/types"
	gl "github.com/rafa-mori/gocrafter/logger"
	"gopkg.in/yaml.v3"
)

// KitManagerImpl implements the KitManager interface
type KitManagerImpl struct {
	config   *types.KitConfig
	kitsPath string
}

// NewKitManager creates a new kit manager instance
func NewKitManager(config *types.KitConfig) (*KitManagerImpl, error) {
	if config == nil {
		// Default configuration
		homeDir, err := os.UserHomeDir()
		if err != nil {
			return nil, fmt.Errorf("failed to get user home directory: %w", err)
		}
		
		config = &types.KitConfig{
			KitsPath:    filepath.Join(homeDir, ".gocrafter", "kits"),
			CachePath:   filepath.Join(homeDir, ".gocrafter", "cache"),
			AutoUpdate:  false,
			MaxCacheAge: 7,
		}
	}

	// Ensure directories exist
	if err := os.MkdirAll(config.KitsPath, 0755); err != nil {
		return nil, fmt.Errorf("failed to create kits directory: %w", err)
	}
	if err := os.MkdirAll(config.CachePath, 0755); err != nil {
		return nil, fmt.Errorf("failed to create cache directory: %w", err)
	}

	return &KitManagerImpl{
		config:   config,
		kitsPath: config.KitsPath,
	}, nil
}

// AddKit adds a new kit from repository URL
func (km *KitManagerImpl) AddKit(repoURL string) error {
	gl.Log("info", fmt.Sprintf("Adding kit from repository: %s", repoURL))

	// Parse repository URL to extract kit name
	kitName := km.extractKitNameFromURL(repoURL)
	if kitName == "" {
		return fmt.Errorf("could not extract kit name from URL: %s", repoURL)
	}

	kitPath := filepath.Join(km.kitsPath, kitName)

	// Check if kit already exists
	if _, err := os.Stat(kitPath); err == nil {
		return fmt.Errorf("kit '%s' already exists. Use update command to update it", kitName)
	}

	// Clone or download the kit
	if err := km.downloadKit(repoURL, kitPath); err != nil {
		return fmt.Errorf("failed to download kit: %w", err)
	}

	// Validate kit structure
	if err := km.ValidateKit(kitPath); err != nil {
		// Clean up on validation failure
		os.RemoveAll(kitPath)
		return fmt.Errorf("kit validation failed: %w", err)
	}

	gl.Log("info", fmt.Sprintf("Kit '%s' added successfully", kitName))
	return nil
}

// RemoveKit removes a kit by name
func (km *KitManagerImpl) RemoveKit(name string) error {
	kitPath := filepath.Join(km.kitsPath, name)
	
	if _, err := os.Stat(kitPath); os.IsNotExist(err) {
		return fmt.Errorf("kit '%s' not found", name)
	}

	if err := os.RemoveAll(kitPath); err != nil {
		return fmt.Errorf("failed to remove kit '%s': %w", name, err)
	}

	gl.Log("info", fmt.Sprintf("Kit '%s' removed successfully", name))
	return nil
}

// ListKits returns all available kits
func (km *KitManagerImpl) ListKits() ([]types.Kit, error) {
	var kits []types.Kit

	entries, err := os.ReadDir(km.kitsPath)
	if err != nil {
		return nil, fmt.Errorf("failed to read kits directory: %w", err)
	}

	for _, entry := range entries {
		if !entry.IsDir() {
			continue
		}

		kitPath := filepath.Join(km.kitsPath, entry.Name())
		kit, err := km.loadKitMetadata(kitPath)
		if err != nil {
			gl.Log("warn", fmt.Sprintf("Failed to load kit metadata for '%s': %v", entry.Name(), err))
			continue
		}

		kit.LocalPath = kitPath
		kits = append(kits, *kit)
	}

	return kits, nil
}

// GetKit returns a specific kit by name
func (km *KitManagerImpl) GetKit(name string) (*types.Kit, error) {
	kitPath := filepath.Join(km.kitsPath, name)
	
	if _, err := os.Stat(kitPath); os.IsNotExist(err) {
		return nil, fmt.Errorf("kit '%s' not found", name)
	}

	kit, err := km.loadKitMetadata(kitPath)
	if err != nil {
		return nil, fmt.Errorf("failed to load kit metadata: %w", err)
	}

	kit.LocalPath = kitPath
	return kit, nil
}

// UpdateKit updates an existing kit
func (km *KitManagerImpl) UpdateKit(name string) error {
	kit, err := km.GetKit(name)
	if err != nil {
		return err
	}

	if kit.Repository == "" {
		return fmt.Errorf("kit '%s' has no repository URL configured", name)
	}

	// Backup current kit
	backupPath := filepath.Join(km.config.CachePath, fmt.Sprintf("%s_backup_%d", name, time.Now().Unix()))
	if err := km.copyDir(kit.LocalPath, backupPath); err != nil {
		return fmt.Errorf("failed to create backup: %w", err)
	}

	// Remove current kit
	if err := os.RemoveAll(kit.LocalPath); err != nil {
		return fmt.Errorf("failed to remove current kit: %w", err)
	}

	// Download updated kit
	if err := km.downloadKit(kit.Repository, kit.LocalPath); err != nil {
		// Restore backup on failure
		km.copyDir(backupPath, kit.LocalPath)
		return fmt.Errorf("failed to download updated kit: %w", err)
	}

	// Validate updated kit
	if err := km.ValidateKit(kit.LocalPath); err != nil {
		// Restore backup on validation failure
		km.copyDir(backupPath, kit.LocalPath)
		return fmt.Errorf("updated kit validation failed: %w", err)
	}

	// Clean up backup
	os.RemoveAll(backupPath)

	gl.Log("info", fmt.Sprintf("Kit '%s' updated successfully", name))
	return nil
}

// ValidateKit validates kit structure and metadata
func (km *KitManagerImpl) ValidateKit(kitPath string) error {
	// Check if metadata.yaml exists
	metadataPath := filepath.Join(kitPath, "metadata.yaml")
	if _, err := os.Stat(metadataPath); os.IsNotExist(err) {
		return fmt.Errorf("metadata.yaml not found in kit")
	}

	// Load and validate metadata
	_, err := km.loadKitMetadata(kitPath)
	if err != nil {
		return fmt.Errorf("invalid metadata: %w", err)
	}

	// Check if templates directory exists
	templatesPath := filepath.Join(kitPath, "templates")
	if info, err := os.Stat(templatesPath); err != nil || !info.IsDir() {
		return fmt.Errorf("templates directory not found in kit")
	}

	return nil
}

// Helper methods

func (km *KitManagerImpl) extractKitNameFromURL(repoURL string) string {
	// Extract kit name from various URL formats
	if strings.Contains(repoURL, "github.com") {
		parts := strings.Split(repoURL, "/")
		if len(parts) >= 2 {
			name := parts[len(parts)-1]
			// Remove .git extension if present
			return strings.TrimSuffix(name, ".git")
		}
	}
	
	// Fallback: use last part of URL path
	parts := strings.Split(strings.TrimSuffix(repoURL, "/"), "/")
	if len(parts) > 0 {
		return strings.TrimSuffix(parts[len(parts)-1], ".git")
	}
	
	return ""
}

func (km *KitManagerImpl) downloadKit(repoURL, targetPath string) error {
	// Check if it's a local path
	if km.isLocalPath(repoURL) {
		return km.copyLocalKit(repoURL, targetPath)
	}
	
	// Try git clone first
	if err := km.gitClone(repoURL, targetPath); err == nil {
		return nil
	}

	// Fallback to HTTP download for archive formats
	return km.httpDownload(repoURL, targetPath)
}

func (km *KitManagerImpl) isLocalPath(path string) bool {
	// Check if it's a local file system path
	if strings.HasPrefix(path, "/") || strings.HasPrefix(path, "./") || strings.HasPrefix(path, "../") {
		return true
	}
	// Check if it's a Windows path
	if len(path) > 1 && path[1] == ':' {
		return true
	}
	return false
}

func (km *KitManagerImpl) copyLocalKit(sourcePath, targetPath string) error {
	// Check if source exists
	if _, err := os.Stat(sourcePath); os.IsNotExist(err) {
		return fmt.Errorf("source path does not exist: %s", sourcePath)
	}

	// Copy the entire directory
	return km.copyDir(sourcePath, targetPath)
}

func (km *KitManagerImpl) gitClone(repoURL, targetPath string) error {
	cmd := exec.Command("git", "clone", repoURL, targetPath)
	if err := cmd.Run(); err != nil {
		return fmt.Errorf("git clone failed: %w", err)
	}
	
	// Remove .git directory to save space
	gitDir := filepath.Join(targetPath, ".git")
	os.RemoveAll(gitDir)
	
	return nil
}

func (km *KitManagerImpl) httpDownload(repoURL, targetPath string) error {
	// This is a simplified implementation for tar.gz archives
	if !strings.HasSuffix(repoURL, ".tar.gz") {
		return fmt.Errorf("unsupported archive format, only .tar.gz supported for HTTP download")
	}

	resp, err := http.Get(repoURL)
	if err != nil {
		return fmt.Errorf("failed to download archive: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("failed to download archive: HTTP %d", resp.StatusCode)
	}

	// Extract tar.gz
	return km.extractTarGz(resp.Body, targetPath)
}

func (km *KitManagerImpl) extractTarGz(src io.Reader, targetPath string) error {
	gzr, err := gzip.NewReader(src)
	if err != nil {
		return err
	}
	defer gzr.Close()

	tr := tar.NewReader(gzr)

	for {
		header, err := tr.Next()
		if err == io.EOF {
			break
		}
		if err != nil {
			return err
		}

		path := filepath.Join(targetPath, header.Name)
		
		switch header.Typeflag {
		case tar.TypeDir:
			if err := os.MkdirAll(path, 0755); err != nil {
				return err
			}
		case tar.TypeReg:
			if err := os.MkdirAll(filepath.Dir(path), 0755); err != nil {
				return err
			}
			
			file, err := os.Create(path)
			if err != nil {
				return err
			}
			
			if _, err := io.Copy(file, tr); err != nil {
				file.Close()
				return err
			}
			file.Close()
		}
	}

	return nil
}

func (km *KitManagerImpl) loadKitMetadata(kitPath string) (*types.Kit, error) {
	metadataPath := filepath.Join(kitPath, "metadata.yaml")
	data, err := os.ReadFile(metadataPath)
	if err != nil {
		return nil, fmt.Errorf("failed to read metadata file: %w", err)
	}

	var kit types.Kit
	if err := yaml.Unmarshal(data, &kit); err != nil {
		return nil, fmt.Errorf("failed to parse metadata YAML: %w", err)
	}

	// Validate required fields
	if kit.Name == "" {
		return nil, fmt.Errorf("kit name is required in metadata")
	}
	if kit.Description == "" {
		return nil, fmt.Errorf("kit description is required in metadata")
	}

	return &kit, nil
}

func (km *KitManagerImpl) copyDir(src, dst string) error {
	return filepath.Walk(src, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}

		relPath, err := filepath.Rel(src, path)
		if err != nil {
			return err
		}

		dstPath := filepath.Join(dst, relPath)

		if info.IsDir() {
			return os.MkdirAll(dstPath, info.Mode())
		}

		srcFile, err := os.Open(path)
		if err != nil {
			return err
		}
		defer srcFile.Close()

		if err := os.MkdirAll(filepath.Dir(dstPath), 0755); err != nil {
			return err
		}

		dstFile, err := os.Create(dstPath)
		if err != nil {
			return err
		}
		defer dstFile.Close()

		_, err = io.Copy(dstFile, srcFile)
		return err
	})
}
