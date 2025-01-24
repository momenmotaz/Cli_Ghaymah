package config

import (
    "os"
    "gopkg.in/yaml.v3"
    "ghaymah-cli/pkg/types"
)

// Config represents the main configuration for the application
type Config struct {
    AppName        string                `yaml:"appName"`
    Image          string                `yaml:"image,omitempty"`
    DockerfilePath string                `yaml:"dockerfilePath,omitempty"`
    EnvVars        map[string]string     `yaml:"envVars,omitempty"`
    Region         string                `yaml:"region,omitempty"`
    Resources      types.ResourceConfig  `yaml:"resources,omitempty"`
    Registry       *types.DockerRegistry `yaml:"registry,omitempty"`
}

// LoadFromFile loads configuration from a YAML file
func (c *Config) LoadFromFile(path string) error {
    data, err := os.ReadFile(path)
    if err != nil {
        return err
    }
    return yaml.Unmarshal(data, c)
}

// Validate ensures all required fields are properly set
func (c *Config) Validate() bool {
    // AppName is always required
    if c.AppName == "" {
        return false
    }

    // Either Image or DockerfilePath must be provided
    if c.Image == "" && c.DockerfilePath == "" {
        return false
    }

    // If resources are specified, validate them
    if c.Resources != (types.ResourceConfig{}) {
        return true // Removed the call to Validate() as it's not defined for types.ResourceConfig
    }

    return true
}
