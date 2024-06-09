package cmd

import (
	"fmt"
	"os"

	"github.com/prateek69/go-productivity/db"
	"github.com/spf13/cobra"
)

var putCmd = &cobra.Command{
	Use:   "put",
	Short: "put a key-val pair",
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) > 2 {
			fmt.Println("Please only add a key and value")
			os.Exit(1)
		}
		err := db.Put(args[0], args[1])
		if err != nil {
			os.Exit(1)
		}
	},
}

var getCmd = &cobra.Command{
	Use:   "get",
	Short: "get a value",
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) > 1 {
			fmt.Println("Please provide only a key!")
			os.Exit(1)
		}
		val, err := db.Get(args[0])
		if err != nil {
			os.Exit(1)
		}
		fmt.Println(val)
	},
}

var delCmd = &cobra.Command{
	Use:   "del",
	Short: "delete a value",
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) > 1 {
			fmt.Println("Please provide only a key!")
			os.Exit(1)
		}
		err := db.Del(args[0])
		if err != nil {
			os.Exit(1)
		}
	},
}

func init() {
	rootCmd.AddCommand(putCmd)
	rootCmd.AddCommand(getCmd)
	rootCmd.AddCommand(delCmd)
}
