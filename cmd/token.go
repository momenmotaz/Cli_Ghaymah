package cmd

import (
    "fmt"
    "github.com/spf13/cobra"
    "ghaymah-cli/pkg/api"
    "ghaymah-cli/pkg/token"
)

func NewTokenCommand(api *api.GhaymahAPI) *cobra.Command {
    cmd := &cobra.Command{
        Use:   "token",
        Short: "Manage API tokens",
        Long:  `Generate, list, and revoke API tokens for your Ghaymah account`,
    }

    cmd.AddCommand(
        newGenerateTokenCommand(api),
        newListTokenCommand(api),
        newRevokeTokenCommand(api),
    )
    return cmd
}

func newGenerateTokenCommand(api *api.GhaymahAPI) *cobra.Command {
    cmd := &cobra.Command{
        Use:   "generate",
        Short: "Generate a new API token",
        RunE: func(cmd *cobra.Command, args []string) error {
            days, _ := cmd.Flags().GetInt("expiry-days")
            
            newToken, err := token.GenerateToken(days)
            if err != nil {
                return fmt.Errorf("failed to generate token: %v", err)
            }

            // TODO: Add API call to save token once API is implemented
            fmt.Printf("Generated new token: %s\n", newToken.Value)
            fmt.Printf("Expires at: %s\n", newToken.ExpiresAt.Format("2006-01-02 15:04:05"))
            return nil
        },
    }

    cmd.Flags().Int("expiry-days", 30, "Token expiration in days")
    return cmd
}

func newListTokenCommand(api *api.GhaymahAPI) *cobra.Command {
    return &cobra.Command{
        Use:   "list",
        Short: "List all active API tokens",
        RunE: func(cmd *cobra.Command, args []string) error {
            // TODO: Add API call to get token list
            // For now, just show a mock response
            fmt.Println("Active API Tokens:")
            fmt.Println("1. token_xyz123... (expires: 2025-02-23)")
            fmt.Println("2. token_abc456... (expires: 2025-03-15)")
            return nil
        },
    }
}

func newRevokeTokenCommand(api *api.GhaymahAPI) *cobra.Command {
    cmd := &cobra.Command{
        Use:   "revoke",
        Short: "Revoke an API token",
        RunE: func(cmd *cobra.Command, args []string) error {
            tokenValue, _ := cmd.Flags().GetString("token")
            
            if tokenValue == "" {
                return fmt.Errorf("token value is required")
            }

            // TODO: Add API call to revoke token
            fmt.Printf("Successfully revoked token: %s\n", tokenValue)
            return nil
        },
    }

    cmd.Flags().String("token", "", "Token value to revoke")
    cmd.MarkFlagRequired("token")

    return cmd
}

func init() {
    rootCmd.AddCommand(NewTokenCommand(ghaymahAPI))
}
