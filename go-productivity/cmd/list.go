package cmd

import (
	"fmt"
	"os"

	"github.com/prateek69/go-productivity/db"
	"github.com/spf13/cobra"
)

var lsCmd = &cobra.Command{
	Use:   "ls",
	Short: "list all the tasks remaining!",
	Run: func(cmd *cobra.Command, args []string) {
		tasks, err := db.GetAllTasks()
		if err != nil {
			os.Exit(1)
		}
		for idx, task := range tasks {
			fmt.Printf("%d | %s\n", idx+1, task.Title)
		}
	},
}

func init() {
	rootCmd.AddCommand(lsCmd)
}
