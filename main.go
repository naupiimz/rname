package main

import (
	"fmt"
	"os"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Usage: rname <filename> [new filename]...")
		return
	}

	err := os.Rename(os.Args[1], os.Args[2])
	if err != nil {
		panic(err)
	}
}
