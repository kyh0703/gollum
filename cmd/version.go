/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

// versionCmd represents the version command
var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "Print the version number of gollum",
	Long:  "",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("version v0.9")
	},
}

func init() {
	rootCmd.AddCommand(versionCmd)
}
