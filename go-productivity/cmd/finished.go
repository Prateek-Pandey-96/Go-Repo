package cmd

import (
	"fmt"
	"os"

	"github.com/prateek69/go-productivity/db"
	"github.com/spf13/cobra"
)

var lsFinishedCmd = &cobra.Command{
	Use:   "lsc",
	Short: "list all the tasks completed within past 24 hours!",
	Run: func(cmd *cobra.Command, args []string) {
		tasks, err := db.GetAllFinishedTasks()
		if err != nil {
			os.Exit(1)
		}
		if len(tasks) == 0 {
			fmt.Println("No finished task in past 24 hours!")
		}
		for idx, task := range tasks {
			fmt.Printf("%d | %s\n", idx+1, task.Title)
		}
	},
}

func init() {
	rootCmd.AddCommand(lsFinishedCmd)
}
