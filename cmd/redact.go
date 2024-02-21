/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

// redactCmd represents the redact command
var redactCmd = &cobra.Command{
	Use:   "redact",
	Short: "Redacts sensitive information from debug bundles",
	Long:  `A tool to automatically redact sensitive information like IP addresses from debug bundles.`,
	Run: func(cmd *cobra.Command, args []string) {
		source, _ := cmd.Flags().GetString("source")
		if source == "" {
			fmt.Println("Source file must be specified")
			return
		}
		err := redactor.Redact(source)
		if err != nil {
			fmt.Println("Error redacting file:", err)
			return
		}
		fmt.Println("File redacted successfully")
	},
}

func init() {
	rootCmd.AddCommand(redactCmd)

	redactCmd.Flags().StringP("source", "s", "", "Source log file to redact")
}
