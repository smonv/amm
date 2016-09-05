package main

import (
	"os"

	"github.com/tthanh/amm"
)

func main() {
	defer os.Exit(1)

	if len(os.Args) < 2 {
		return
	}

	args := os.Args[1:]

	switch args[0] {
	case "init":
		amm.Init()
	case "add":
		if len(args) < 2 {
			return
		}
		amm.Add(args[1:])
	default:
		return
	}
}
