package main

import (
	"github.com/spf13/cobra"
)

func UserCmd() *cobra.Command {
	userCmd := &cobra.Command{
		Use:   "user",
		Short: "Manage users",
	}

	// Subcommands for the "user" command
	userCmd.AddCommand(userAddCmd(), userDeleteCmd(), userUpdateCmd())
	return userCmd
}

func userAddCmd() *cobra.Command {
	var username, password string
	cmd := &cobra.Command{
		Use:   "add",
		Short: "Add a new user",
		Run: func(cmd *cobra.Command, args []string) {
			//fmt.Printf("Adding user: %s with password: %s\n", username, password)
			addUser(cfg.URL, cfg.RootUsername, cfg.RootPassword, username, password)
		},
	}
	cmd.Flags().StringVarP(&username, "username", "u", "", "Username")
	cmd.Flags().StringVarP(&password, "password", "p", "", "Password")
	cmd.MarkFlagRequired("username")
	cmd.MarkFlagRequired("password")
	return cmd
}

func userDeleteCmd() *cobra.Command {
	var username string
	cmd := &cobra.Command{
		Use:   "delete",
		Short: "Delete a user",
		Run: func(cmd *cobra.Command, args []string) {
			//fmt.Printf("Deleting user: %s\n", username)
			deleteUser(cfg.URL, cfg.RootUsername, cfg.RootPassword, username)
		},
	}
	cmd.Flags().StringVarP(&username, "username", "u", "", "Username")
	cmd.MarkFlagRequired("username")
	return cmd
}

func userUpdateCmd() *cobra.Command {
	var username, password string
	cmd := &cobra.Command{
		Use:   "update",
		Short: "Update user credentials",
		Run: func(cmd *cobra.Command, args []string) {
			//fmt.Printf("Updating user: %s with new password: %s\n", username, password)
			updateUserCredentials(cfg.URL, cfg.Username, cfg.Password, username, password)
		},
	}
	cmd.Flags().StringVarP(&username, "username", "u", "", "Username")
	cmd.Flags().StringVarP(&password, "password", "p", "", "New Password")
	cmd.MarkFlagRequired("username")
	cmd.MarkFlagRequired("password")
	return cmd
}
