package cmd

import (
	"fmt"
	"os"
	"path"
	"strings"

	"github.com/aeoss/fusionauth-cli/pkg/batcher"
	"github.com/aeoss/fusionauth-cli/pkg/client"
	"github.com/spf13/cobra"

	"github.com/spf13/viper"
)

var batch *batcher.Client

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "factl",
	Short: "A CLI for FusionAuth management",
	Long: `This CLI provides easy management utilities for FusionAuth
environments to allow management of FusionAuth using a GitOps approach.
`,
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func init() {
	cobra.OnInitialize(initConfig)
}

// initConfig reads in config file and ENV variables if set.
func initConfig() {
	viper.AddConfigPath(path.Join(".factl"))
	viper.SetConfigName("config")

	viper.AutomaticEnv() // read in environment variables that match

	// If a config file is found, read it in.
	if err := viper.ReadInConfig(); err != nil {
		fmt.Println("No config file could be read. Assuming you are using environment config and will continue...")
	}

	var currentEnv = viper.GetString("currentEnvironment")
	if currentEnv == "" {
		currentEnv = "default"
	}

	batch = batcher.Init(".", client.Init(&client.Context{
		URL: viper.GetString(strings.Join([]string{"environment", currentEnv, "url"}, ".")),
		Key: viper.GetString(strings.Join([]string{"environment", currentEnv, "key"}, ".")),
		JWT: viper.GetBool(strings.Join([]string{"environment", currentEnv, "jwt"}, ".")),
	}))
}
