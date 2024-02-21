package cmd

import (
	"debug-redactor/pkg/redactor" // Adjust the import path as necessary
	"fmt"
	"github.com/spf13/cobra"
	"os"
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

		// Call the function and handle its returned values properly
		outputFilePath, err := redactor.ProcessAndRepackageTarGz(source)
		if err != nil {
			fmt.Printf("Error processing file: %s\n", err)
			return
		}

		// Use the outputFilePath in the success message to indicate where the file was processed
		fmt.Printf("File processed successfully. Output file: %s\n", outputFilePath)
	},
}

func init() {
	// Register flags and other initializations here
	rootCmd.Flags().StringP("source", "s", "", "Path to the source tar.gz file to process (required)")
	rootCmd.MarkFlagRequired("source")
}

// Execute adds all child commands to the root command sets flags appropriately.
func Execute() {
	if err := rootCmd.Execute(); err != nil {
		os.Exit(1)
	}
}
