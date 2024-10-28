package main

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "mycli",
	Short: "My cli application",
	Long:  "A longer description of my Cli application.",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Welcome to my cli")
	},
}

func main() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
