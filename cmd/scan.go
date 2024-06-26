/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"os"
	"strconv"

	"github.com/alex1988m/go-pscan/scan"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

// scanCmd represents the scan command
var scanCmd = &cobra.Command{
	Use:          "scan",
	Short:        "scan hosts ports",
	SilenceUsage: true,
	DisableAutoGenTag: true,
	RunE: func(cmd *cobra.Command, args []string) error {
		// validation
		hostsfile := viper.GetString("hosts-file")
		rawPorts := viper.GetString("ports")
		rawRange := viper.GetString("range")
		filter := viper.GetString("filter")
		rawTimeout := viper.GetString("timeout")
		timeout := 1000
		if rawTimeout != "" {
			cfgTimeout, err := strconv.Atoi(rawTimeout)
			if err != nil {
				return err
			}
			if cfgTimeout < 1 {
				return fmt.Errorf("timeout must be greater than 0")
			}
			timeout = cfgTimeout
		}
		if filter != "" && filter != "open" && filter != "closed" {
			return fmt.Errorf("invalid filter: %s, accepted values: open, closed", filter)
		}
		hl := &scan.HostsList{Filename: hostsfile, W: os.Stdout}
		if err := hl.Load(); err != nil {
			return err
		}
		ports, err := scan.ToPortList(rawPorts, rawRange)
		if err != nil {
			return err
		}
		ps := &scan.PortScanner{Hosts: hl.Hosts, Ports: ports, W: os.Stdout, Filter: filter, Timeout: timeout}
	
		if err := ps.ValidateHosts(); err != nil {
			return err
		}
		ps.ScanPorts()
		ps.SortResults()
		if err := ps.PrintResults(); err != nil {
			return err
		}
		return nil
	},
}

func init() {
	rootCmd.AddCommand(scanCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// scanCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	scanCmd.PersistentFlags().StringP("ports", "p", "22,80,443", "ports to scan within hosts")
	scanCmd.PersistentFlags().StringP("range", "r", "", "port range to scan within hosts")
	scanCmd.PersistentFlags().StringP("filter", "f", "", "filter open or closed ports")
	scanCmd.PersistentFlags().StringP("timeout", "t", "", "timeout in milliseconds")
	viper.BindPFlag("ports", scanCmd.PersistentFlags().Lookup("ports"))
	viper.BindPFlag("filter", scanCmd.PersistentFlags().Lookup("filter"))
	viper.BindPFlag("range", scanCmd.PersistentFlags().Lookup("range"))
	viper.BindPFlag("timeout", scanCmd.PersistentFlags().Lookup("timeout"))

}
