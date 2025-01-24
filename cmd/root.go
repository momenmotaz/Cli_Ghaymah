package cmd

import (
    "fmt"
    "os"

    "github.com/spf13/cobra"
    "ghaymah-cli/pkg/api"
    "ghaymah-cli/pkg/config"
)

var (
    rootCmd = &cobra.Command{
        Use:   "ghaymah",
        Short: "Ghaymah CLI - Cloud Platform Management Tool",
        Long: `Ghaymah CLI is a command line tool for managing applications
and resources on the Ghaymah Cloud Platform.`,
    }
    ghaymahAPI *api.GhaymahAPI
)

func init() {
    // Get API configuration from environment variables
    apiURL, apiToken, err := config.GetAPIConfig()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error: %v\n", err)
        fmt.Fprintf(os.Stderr, "\nPlease set the following environment variables:\n")
        fmt.Fprintf(os.Stderr, "  %s: The URL of the Ghaymah Cloud API\n", config.APIURLEnvVar)
        fmt.Fprintf(os.Stderr, "  %s: Your Ghaymah Cloud API token\n", config.APITokenEnvVar)
        os.Exit(1)
    }

    // Create API client
    ghaymahAPI = api.NewGhaymahAPI(apiURL, apiToken)

    // Add commands
    rootCmd.AddCommand(
        NewDeployCommand(ghaymahAPI),
        NewStatusCommand(ghaymahAPI),
        NewLogsCommand(ghaymahAPI),
        NewRegistryCommand(ghaymahAPI),
        NewTokenCommand(ghaymahAPI),
    )
}

// Execute adds all child commands to the root command and sets flags appropriately.
func Execute() {
    if err := rootCmd.Execute(); err != nil {
        fmt.Println(err)
        os.Exit(1)
    }
}
