/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"os"

	"github.com/alex1988m/go-pscan/scan"
	"github.com/spf13/cobra"
)

// deleteCmd represents the delete command
var deleteCmd = &cobra.Command{
	Use:   "delete",
	Short: "delete hosts from the hosts list",
	Args:  cobra.MinimumNArgs(1),
	SilenceUsage: true,
	RunE: func(cmd *cobra.Command, args []string) error {
		hostsfile, err := cmd.Flags().GetString("hosts-file")
		if err != nil {
			return err
		}
		hl := &scan.HostsList{Filename: hostsfile, W: os.Stdout}
		if err := hl.Remove(args); err != nil {
			return err
		}
		return nil
	},
}

func init() {
	hostsCmd.AddCommand(deleteCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// deleteCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// deleteCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
