package main

import (
	"github.com/spf13/cobra"
)

func SecretCmd() *cobra.Command {
	secretCmd := &cobra.Command{
		Use:   "secret",
		Short: "Manage secrets",
	}

	secretCmd.AddCommand(secretAddCmd(), secretDeleteCmd(), secretUpdateCmd(), secretGetCmd())
	return secretCmd
}

func secretAddCmd() *cobra.Command {
	var secretID, secret string
	cmd := &cobra.Command{
		Use:   "add",
		Short: "Add a new secret",
		Run: func(cmd *cobra.Command, args []string) {
			addSecret(cfg.URL, cfg.Username, cfg.Password, secretID, secret, cfg.PublicKey)
		},
	}
	cmd.Flags().StringVarP(&secretID, "secret-id", "i", "", "ID of the secret")
	cmd.Flags().StringVarP(&secret, "secret", "s", "", "The secret data")
	cmd.MarkFlagRequired("secret-id")
	cmd.MarkFlagRequired("secret")
	return cmd
}

func secretDeleteCmd() *cobra.Command {
	var secretID string
	cmd := &cobra.Command{
		Use:   "delete",
		Short: "Delete a secret",
		Run: func(cmd *cobra.Command, args []string) {
			deleteSecret(cfg.URL, cfg.Username, cfg.Password, secretID)
		},
	}
	cmd.Flags().StringVarP(&secretID, "secret-id", "i", "", "ID of the secret")
	cmd.MarkFlagRequired("secret-id")
	return cmd
}

func secretUpdateCmd() *cobra.Command {
	var secretID, secret string
	cmd := &cobra.Command{
		Use:   "update",
		Short: "Update an existing secret",
		Run: func(cmd *cobra.Command, args []string) {
			updateSecret(cfg.URL, cfg.Username, cfg.Password, secretID, secret, cfg.PublicKey)
		},
	}
	cmd.Flags().StringVarP(&secretID, "secret-id", "i", "", "ID of the secret")
	cmd.Flags().StringVarP(&secret, "secret", "s", "", "New secret data")
	cmd.MarkFlagRequired("secret-id")
	cmd.MarkFlagRequired("secret")
	return cmd
}

func secretGetCmd() *cobra.Command {
	var secretID string
	cmd := &cobra.Command{
		Use:   "get",
		Short: "Get a secret",
		Run: func(cmd *cobra.Command, args []string) {
			getSecret(cfg.URL, cfg.Username, cfg.Password, secretID, cfg.PrivateKey)
		},
	}
	cmd.Flags().StringVarP(&secretID, "secret-id", "i", "", "ID of the secret")
	cmd.MarkFlagRequired("secret-id")
	return cmd
}
