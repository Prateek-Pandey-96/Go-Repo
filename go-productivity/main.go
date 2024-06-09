package main

import (
	"os"

	"github.com/mitchellh/go-homedir"
	"github.com/prateek69/go-productivity/cmd"
	"github.com/prateek69/go-productivity/db"
)

func main() {
	path, _ := homedir.Dir()
	handleError(db.Init(path))
	handleError(cmd.Execute())
}

func handleError(err error) {
	if err != nil {
		os.Exit(1)
	}
}
