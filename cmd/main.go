package main

import (
	"github.com/MinseokOh/data-warehouse/cmd/command"
	"os"
)

func main() {
	if err := command.NewRootCmd().Execute(); err != nil {
		panic(err)
		os.Exit(1)
	}
}
