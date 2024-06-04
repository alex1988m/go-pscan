/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
	"github.com/spf13/cobra/doc"
	"github.com/spf13/viper"
)

// docsCmd represents the docs command
var docsCmd = &cobra.Command{
	Use:   "docs",
	Short: "Generate documentation for your command",
	RunE: func(cmd *cobra.Command, args []string) error {
		dir := viper.GetString("dir")
		var err error
		if len(dir) > 0 {
			if err := os.RemoveAll(dir); err != nil {
				return err
			}
			if err := os.Mkdir(dir, 0755); err != nil {
				return err
			}
		} else {
			dir, err = os.MkdirTemp("", "pscan")
			if err != nil {
				return err
			}
		}

		if err := doc.GenMarkdownTree(rootCmd, dir); err != nil {
			return err
		}
		fmt.Printf("Documentation generated in %s\n", dir)
		return nil
	},
}

func init() {
	rootCmd.AddCommand(docsCmd)
	docsCmd.Flags().StringP("dir", "d", "", "destination directory for docs")
	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// docsCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// docsCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
