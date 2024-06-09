package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "ptv",
	Short: "productive is a cli tool for daily tasks!",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Print("I am cobra")
	},
}

func Execute() error {
	if err := rootCmd.Execute(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		return err
	}
	return nil
}
