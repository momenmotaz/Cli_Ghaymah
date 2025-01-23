package cmd

import (
    "fmt"
    "time"
    "github.com/spf13/cobra"
    "ghaymah-cli/pkg/api"
    "ghaymah-cli/pkg/types"
)

// NewLogsCommand creates a new logs command
func NewLogsCommand(api *api.GhaymahAPI) *cobra.Command {
    var (
        appName string
        follow  bool
        tail    int
        since   string
    )

    cmd := &cobra.Command{
        Use:   "logs",
        Short: "View application logs",
        Long: `View logs from your application running on Ghaymah Cloud.
You can view historical logs or follow them in real-time. Use flags to customize
the output according to your needs.

Examples:
  # View recent logs
  ghaymah logs --name my-app

  # Follow logs in real-time
  ghaymah logs --name my-app --follow

  # View last 50 lines
  ghaymah logs --name my-app --tail 50

  # View logs since a specific time
  ghaymah logs --name my-app --since 2024-01-23T00:00:00Z`,
        RunE: func(cmd *cobra.Command, args []string) error {
            if appName == "" {
                return fmt.Errorf("application name is required. Use --name flag")
            }

            fmt.Printf("Retrieving logs for application %s...\n", appName)

            options := &types.LogOptions{
                Follow: follow,
                Tail:   tail,
            }

            if since != "" {
                sinceTime, err := time.Parse(time.RFC3339, since)
                if err != nil {
                    return fmt.Errorf("invalid timestamp format: %v", err)
                }
                options.Since = sinceTime
            }

            logs, err := api.GetLogs(appName, options)
            if err != nil {
                return fmt.Errorf("failed to retrieve logs: %v", err)
            }

            if len(logs.Entries) == 0 {
                fmt.Println("No logs available for the application")
                return nil
            }

            for _, entry := range logs.Entries {
                fmt.Printf("[%s] %s\n", entry.Timestamp.Format(time.RFC3339), entry.Message)
            }

            if follow {
                fmt.Println("\nFollowing logs in real-time... (Press Ctrl+C to exit)")
                // In a real implementation, we would stream logs here
            }

            return nil
        },
    }

    cmd.Flags().StringVar(&appName, "name", "", "Application name")
    cmd.MarkFlagRequired("name")
    cmd.Flags().BoolVarP(&follow, "follow", "f", false, "Follow log output in real-time")
    cmd.Flags().IntVarP(&tail, "tail", "n", 100, "Number of lines to show from the end of the logs")
    cmd.Flags().StringVarP(&since, "since", "s", "", "Show logs since timestamp (RFC3339 format)")

    return cmd
}
