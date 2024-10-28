package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var primeCmd = &cobra.Command{
	Use:   "Check [prime number]",
	Short: "Check number",
	Long:  "Check new task to your task list",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("check prime number")
	},
}
