package main

import (
	"os"

	"github.com/tthanh/amm"
)

func main() {
	defer os.Exit(1)

	args := os.Args[1:]

	if len(args) < 1 {
		return
	}

	switch args[0] {
	case "init":
		amm.Init()
	default:
		return
	}
}
