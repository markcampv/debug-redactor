package cmd

import (
	"debug-redactor/pkg/redactor"
	"fmt"
	"github.com/spf13/cobra"
)

// redactCmd represents the redact command
var redactCmd = &cobra.Command{
	Use:   "redact",
	Short: "Redacts sensitive information from debug bundles",
	Long: `A tool to automatically redact sensitive information like IP addresses from debug bundles.
Use this command to specify the source tar.gz file that you want to process.`,
	Run: func(cmd *cobra.Command, args []string) {
		source, _ := cmd.Flags().GetString("source")
		if source == "" {
			fmt.Println("Source file must be specified with --source flag")
			return
		}
		err := redactor.ProcessAndRepackageTarGz(source)
		if err != nil {
			fmt.Println("Error redacting file:", err)
			return
		}
		redactedTarGzPath := redactor.GenerateRedactedPath(source)
		fmt.Printf("File redacted successfully. Redacted file is located at: %s\n", redactedTarGzPath)
	},
}

func init() {
	rootCmd.AddCommand(redactCmd)
	redactCmd.Flags().StringP("source", "s", "", "Path to the source tar.gz file to redact (required)")
	redactCmd.MarkFlagRequired("source") // This makes the source flag required
}
