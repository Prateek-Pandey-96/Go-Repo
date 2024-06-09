package cmd

import (
	"os"

	"github.com/prateek69/go-productivity/db"
	"github.com/spf13/cobra"
)

var addCmd = &cobra.Command{
	Use:   "add",
	Short: "add a task to do!",
	Run: func(cmd *cobra.Command, args []string) {
		task := ""
		for idx, arg := range args {
			task += arg
			if idx != len(args) {
				task += " "
			}
		}

		err := db.AddTask(task)
		if err != nil {
			os.Exit(1)
		}
	},
}

func init() {
	rootCmd.AddCommand(addCmd)
}
