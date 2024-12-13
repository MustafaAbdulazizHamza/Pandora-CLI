package main

import (
	"fmt"
	"github.com/spf13/cobra"
	"os"
)

type Config struct {
	URL          string `yaml:"url"`
	RootUsername string `yaml:"root_username"`
	RootPassword string `yaml:"root_password"`
	Username     string `yaml:"username"`
	Password     string `yaml:"password"`
	PrivateKey   string `yaml:"private_key"`
	PublicKey    string `yaml:"public_key"`
}

var (
	cfg Config
)

func main() {

	rootCmd := &cobra.Command{Use: "pandora-client"}
	loc := getExecutableLocation()
	loadConfig(fmt.Sprintf("%s/config.yaml", loc))
	// Add "secret" and "user" commands to the root command
	rootCmd.AddCommand(SecretCmd(), UserCmd())

	// Execute the root command
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

}
