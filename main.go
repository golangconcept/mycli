package main

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

// var primeCmd = &cobra.Command{
// 	Use:   "Check [prime number]",
// 	Short: "Check number",
// 	Long:  "Check new task to your task list",
// 	Args:  cobra.ExactArgs(1),
// 	Run: func(cmd *cobra.Command, args []string) {
// 		fmt.Println("check prime number")
// 	},
// }

var rootCmd = &cobra.Command{
	Use:   "mycli",
	Short: "My cli application",
	Long:  "A longer description of my Cli application.",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Welcome to my cli")
	},
}

func main() {
	rootCmd.AddCommand(cmd.primeCmd)
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
