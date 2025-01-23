package main

import (
    "fmt"
    "os"

    "github.com/spf13/cobra"
    "ghaymah-cli/cmd"
    "ghaymah-cli/pkg/api"
    "ghaymah-cli/pkg/config"
)

func main() {
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
    api := api.NewGhaymahAPI(apiURL, apiToken)

    // Create root command
    rootCmd := &cobra.Command{
        Use:   "ghaymah",
        Short: "Command line interface for Ghaymah Cloud",
        Long: `Ghaymah CLI is a powerful tool for managing your applications on Ghaymah Cloud.
Deploy, monitor, and manage your applications with simple commands.`,
    }

    // Add commands
    rootCmd.AddCommand(
        cmd.NewDeployCommand(api),
        cmd.NewStatusCommand(api),
        cmd.NewLogsCommand(api),
    )

    // Execute
    if err := rootCmd.Execute(); err != nil {
        fmt.Fprintf(os.Stderr, "Error: %v\n", err)
        os.Exit(1)
    }
}
