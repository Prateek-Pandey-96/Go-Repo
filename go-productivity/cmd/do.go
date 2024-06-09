package cmd

import (
	"fmt"
	"os"
	"strconv"

	"github.com/prateek69/go-productivity/db"
	"github.com/spf13/cobra"
)

var doCmd = &cobra.Command{
	Use:   "do",
	Short: "complete a task!",
	Run: func(cmd *cobra.Command, args []string) {
		tasks, err := db.GetAllTasks()
		if err != nil {
			os.Exit(1)
		}

		if len(args) > 1 {
			fmt.Println("Please finish one task at a time!")
		}

		idxToId := make(map[int]uint64, len(tasks)+1)
		for idx, task := range tasks {
			idxToId[idx+1] = task.Id
		}

		entered_id, err := strconv.Atoi(args[0])
		if err != nil {
			fmt.Println("Please enter a valid id!")
		}

		finish_task := tasks[entered_id-1].Title
		err = db.FinishTask(finish_task)
		if err != nil {
			os.Exit(1)
		}

		err = db.DeleteTask(idxToId[entered_id])
		if err != nil {
			os.Exit(1)
		}
	},
}

func init() {
	rootCmd.AddCommand(doCmd)
}
