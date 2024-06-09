package cmd

import (
	"fmt"
	"os"
	"strconv"

	"github.com/prateek69/go-productivity/db"
	"github.com/spf13/cobra"
)

var rmCmd = &cobra.Command{
	Use:   "rm",
	Short: "remove a task from list!",
	Run: func(cmd *cobra.Command, args []string) {
		tasks, err := db.GetAllTasks()
		if err != nil {
			os.Exit(1)
		}

		if len(args) > 1 {
			fmt.Println("Please remove one task at a time!")
		}

		idxToId := make(map[int]uint64, len(tasks)+1)
		for idx, task := range tasks {
			idxToId[idx+1] = task.Id
		}

		task_id, err := strconv.Atoi(args[0])
		if err != nil {
			fmt.Println("Please enter a valid id!")
		}

		err = db.DeleteTask(idxToId[task_id])
		if err != nil {
			os.Exit(1)
		}
	},
}

func init() {
	rootCmd.AddCommand(rmCmd)
}
