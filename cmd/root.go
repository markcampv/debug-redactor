package cmd

import (
	"debug-redactor/pkg/redactor" // Adjust the import path as necessary
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "debug-redactor",
	Short: "A brief description of your application",
	Long: `A longer description that spans multiple lines and likely contains
examples and usage of using your application. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	// This function is called when rootCmd executes
	Run: func(cmd *cobra.Command, args []string) {
		source, err := cmd.Flags().GetString("source")
		if err != nil || source == "" {
			fmt.Println("Error: Source file must be specified with --source flag")
			return
		}
		if err := redactor.ProcessAndRepackageTarGz(source); err != nil {
			fmt.Printf("Error processing file: %s\n", err)
			return
		}
		fmt.Println("File processed successfully")
	},
}

func init() {
	// Define your flags and configuration settings here.
	// For example, you can define the "source" flag for the root command
	rootCmd.Flags().StringP("source", "s", "", "Path to the source tar.gz file to process (required)")
	rootCmd.MarkFlagRequired("source") // Mark the "source" flag as required
}

// Execute adds all child commands to the root command sets flags appropriately.
func Execute() {
	if err := rootCmd.Execute(); err != nil {
		os.Exit(1)
	}
}
