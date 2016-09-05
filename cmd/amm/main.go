package main

import (
	"fmt"
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
		err := amm.Init()
		if err != nil {
			fmt.Println(err)
			return
		}
		fmt.Println("metadata.json created")
	}
}
