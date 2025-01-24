package cmd

import (
    "fmt"
    "github.com/spf13/cobra"
    "ghaymah-cli/pkg/api"
    "ghaymah-cli/pkg/types"
)

func NewRegistryCommand(api *api.GhaymahAPI) *cobra.Command {
    cmd := &cobra.Command{
        Use:   "registry",
        Short: "Manage Docker registry credentials",
        Long:  `Add, update, or remove Docker registry credentials for your Ghaymah account`,
    }

    cmd.AddCommand(
        newAddRegistryCommand(api),
        newListRegistryCommand(api),
        newRemoveRegistryCommand(api),
    )
    return cmd
}

func newAddRegistryCommand(api *api.GhaymahAPI) *cobra.Command {
    cmd := &cobra.Command{
        Use:   "add",
        Short: "Add new Docker registry credentials",
        RunE: func(cmd *cobra.Command, args []string) error {
            registryURL, _ := cmd.Flags().GetString("url")
            username, _ := cmd.Flags().GetString("username")
            password, _ := cmd.Flags().GetString("password")

            registry := &types.DockerRegistry{
                RegistryURL: registryURL,
                Username:    username,
                Password:    password,
            }

            if !registry.Validate() {
                return fmt.Errorf("invalid registry configuration: all fields are required")
            }

            // TODO: Add API call to save registry credentials once API is implemented
            fmt.Printf("Successfully added registry credentials for %s\n", registryURL)
            return nil
        },
    }

    cmd.Flags().String("url", "", "Registry URL (e.g., docker.io)")
    cmd.Flags().String("username", "", "Registry username")
    cmd.Flags().String("password", "", "Registry password")

    cmd.MarkFlagRequired("url")
    cmd.MarkFlagRequired("username")
    cmd.MarkFlagRequired("password")

    return cmd
}

func newListRegistryCommand(api *api.GhaymahAPI) *cobra.Command {
    return &cobra.Command{
        Use:   "list",
        Short: "List all registered Docker registries",
        RunE: func(cmd *cobra.Command, args []string) error {
            // TODO: Add API call to get registry list
            // For now, just show a mock response
            fmt.Println("Registered Docker Registries:")
            fmt.Println("1. docker.io (username: user1)")
            fmt.Println("2. ghcr.io (username: user2)")
            return nil
        },
    }
}

func newRemoveRegistryCommand(api *api.GhaymahAPI) *cobra.Command {
    cmd := &cobra.Command{
        Use:   "remove",
        Short: "Remove Docker registry credentials",
        RunE: func(cmd *cobra.Command, args []string) error {
            registryURL, _ := cmd.Flags().GetString("url")
            
            if registryURL == "" {
                return fmt.Errorf("registry URL is required")
            }

            // TODO: Add API call to remove registry
            fmt.Printf("Successfully removed registry credentials for %s\n", registryURL)
            return nil
        },
    }

    cmd.Flags().String("url", "", "Registry URL to remove")
    cmd.MarkFlagRequired("url")

    return cmd
}

func init() {
    rootCmd.AddCommand(NewRegistryCommand(ghaymahAPI))
}
