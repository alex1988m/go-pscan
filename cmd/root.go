/*
Copyright © 2024
*/
package cmd

import (
	"fmt"
	"os"
	"strings"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var cfgFile string = ""

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "pscan",
	DisableAutoGenTag: true,
	Short: "Fast TCP port scanner",
	Long: `pscan - short for Port Scanner - executes TCP port scan on a list of hosts.
	pScan allows you to add, list, and delete hosts from the list.
	pScan executes a port scan on specified TCP ports. You can customize the
	target ports using a command line flag.`,
	Version: "0.0.1",
	// Uncomment the following line if your bare application
	// has an action associated with it:
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	cobra.OnInitialize(initConfig)
	// Here you will define your flags and configuration settings.
	// Cobra supports persistent flags, which, if defined here,
	// will be global for your application.

	rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.pscan.yaml)")
	rootCmd.PersistentFlags().StringP("hosts-file", "F", "pscan.hosts", "file to store hosts")

	replacer := strings.NewReplacer("-", "_")
	viper.SetEnvKeyReplacer(replacer)
	viper.SetEnvPrefix("PSCAN")
	viper.BindPFlag("hosts-file", rootCmd.PersistentFlags().Lookup("hosts-file"))

	versionTemplate := `{{printf "%s: %s - version %s\n" .Name .Short .Version}}`
	rootCmd.SetVersionTemplate(versionTemplate)
}

func initConfig() {
	viper.AddConfigPath(".")
	viper.SetConfigName(".pscan")
	viper.AutomaticEnv()
	if err := viper.ReadInConfig(); err == nil {
		fmt.Println("Using config file:", viper.ConfigFileUsed())
	}
}
