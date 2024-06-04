/*
Copyright © 2024
*/
package cmd

import (
	"os"

	"github.com/alex1988m/go-pscan/scan"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

// listCmd represents the list command
var listCmd = &cobra.Command{
	Use:     "list",
	Aliases: []string{"l"},
	DisableAutoGenTag: true,
	Short:   "used to list hosts from hosts file",
	RunE: func(cmd *cobra.Command, args []string) error {
		hostsfile := viper.GetString("hosts-file")
		hl := &scan.HostsList{Filename: hostsfile, W: os.Stdout}
		if err := hl.Print(); err != nil {
			return err
		}
		return nil
	},
}

func init() {
	hostsCmd.AddCommand(listCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// listCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// listCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
