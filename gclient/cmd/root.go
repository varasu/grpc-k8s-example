package cmd

import (
	"os"

	"github.com/spf13/cobra"
)

var url string

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "gclient",
	Short: "Client for grpc k8s example",
	Long:  `Execute commands to gservice grpc server.`,
}

func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	rootCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
	rootCmd.PersistentFlags().StringVar(&url, "url", "localhost:8080", "gservice url")
}
