package cmd

import (
    "fmt"
    "github.com/spf13/cobra"
    "ghaymah-cli/pkg/api"
)

// NewStatusCommand creates a new status command
func NewStatusCommand(api *api.GhaymahAPI) *cobra.Command {
    var appName string

    cmd := &cobra.Command{
        Use:   "status",
        Short: "Check application status",
        Long: `View the current status of your application on Ghaymah Cloud.
This includes deployment status, resource usage, and health metrics.

Example:
  ghaymah status --name my-app`,
        RunE: func(cmd *cobra.Command, args []string) error {
            if appName == "" {
                return fmt.Errorf("application name is required. Use --name flag")
            }

            fmt.Printf("Checking status for application %s...\n", appName)

            status, err := api.GetStatus(appName)
            if err != nil {
                return fmt.Errorf("failed to get status: %v", err)
            }

            fmt.Printf("\nStatus: %s\n", status.State)
            fmt.Printf("Last Deployment: %s\n\n", status.LastDeployment.Format("2006-01-02 15:04:05"))

            fmt.Println("Resource Usage:")
            fmt.Printf("  CPU: %.2f%%\n", status.Resources.CPUUsage)
            fmt.Printf("  Memory: %.2f%%\n", status.Resources.MemoryUsage)
            fmt.Printf("  Storage: %.2f%%\n", status.Resources.StorageUsage)

            return nil
        },
    }

    cmd.Flags().StringVar(&appName, "name", "", "Application name")
    cmd.MarkFlagRequired("name")

    return cmd
}
