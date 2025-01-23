package cmd

import (
    "fmt"
    "strings"
    "github.com/spf13/cobra"
    "ghaymah-cli/pkg/api"
    "ghaymah-cli/pkg/config"
)

var (
    configFile string
    imageName  string
    appName    string
)

// DeployCommand handles application deployment
type DeployCommand struct {
    config *config.Config
    api    api.GhaymahAPI
}

// NewDeployCommand creates a new deploy command
func NewDeployCommand(api *api.GhaymahAPI) *cobra.Command {
    deployCmd := &DeployCommand{
        api: *api,
    }

    cmd := &cobra.Command{
        Use:   "deploy",
        Short: "Deploy an application to Ghaymah Cloud",
        Long: `Deploy your application to Ghaymah Cloud platform.
You can deploy either using a configuration file or directly using flags.

Examples:
  # Deploy using config file
  ghaymah deploy -c config.yaml

  # Deploy using image
  ghaymah deploy --image username/app:tag

  # Deploy using image and custom name
  ghaymah deploy --image username/app:tag --name my-app`,
        RunE: func(cmd *cobra.Command, args []string) error {
            var cfg *config.Config

            // If image is provided, create config from flags
            if imageName != "" {
                cfg = &config.Config{
                    AppName: appName,
                    Image:   imageName,
                }
                if cfg.AppName == "" {
                    // If name not provided, use image name without tag
                    cfg.AppName = getAppNameFromImage(imageName)
                }
            } else {
                // Load from config file
                cfg = &config.Config{}
                if err := cfg.LoadFromFile(configFile); err != nil {
                    return fmt.Errorf("failed to load config: %v", err)
                }
            }

            deployCmd.config = cfg

            return deployCmd.Execute()
        },
    }

    // Add flags
    cmd.Flags().StringVarP(&configFile, "config", "c", "ghaymah.yaml", "path to configuration file")
    cmd.Flags().StringVar(&imageName, "image", "", "Docker image to deploy (e.g., username/app:tag)")
    cmd.Flags().StringVar(&appName, "name", "", "Application name (optional when using --image)")

    return cmd
}

// Execute runs the deployment process
func (d *DeployCommand) Execute() error {
    if err := d.validateConfig(); err != nil {
        return err
    }

    fmt.Printf("Starting deployment of %s...\n", d.config.AppName)

    // Deploy application
    resp, err := d.api.Deploy(d.config)
    if err != nil {
        return fmt.Errorf("deployment failed: %v", err)
    }

    fmt.Printf("Successfully deployed! Application ID: %s\n", resp.AppID)
    return nil
}

// validateConfig ensures all required configuration is present
func (d *DeployCommand) validateConfig() error {
    if d.config == nil {
        return fmt.Errorf("configuration is required")
    }
    if !d.config.Validate() {
        return fmt.Errorf("invalid configuration")
    }
    return nil
}

// getAppNameFromImage extracts app name from image (e.g., "myapp" from "username/myapp:v1")
func getAppNameFromImage(image string) string {
    // Remove tag if exists
    name := strings.Split(image, ":")[0]
    
    // Remove registry/username if exists
    parts := strings.Split(name, "/")
    return parts[len(parts)-1]
}
