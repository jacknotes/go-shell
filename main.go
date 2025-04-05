package main

import (
	"fmt"

	"github.com/jacknotes/go-shell.git/cmd"
)

func main() {
	if err := cmd.RootCmd.Execute(); err != nil {
		fmt.Println(err)
	}
}
