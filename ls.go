package main

import (
	"fmt"
	"os"
)

func main() {
	var dir string
	var err error
	if len(os.Args) < 2 {
		dir, err = os.Getwd()
		if err != nil {
			panic("Couldn't open current working directory")
		}
	} else {
		dir = os.Args[1]
	}

	file, err := os.Open(dir)
	if err != nil {
		panic("Couldn't open directory. " + err.Error())
	}

	names, err := file.Readdirnames(0)
	if err != nil {
		panic(err.Error())
	}

	for i, name := range names {
		if i < len(names) {
			fmt.Print(" ")
		}

		fmt.Print(name)
	}

	fmt.Print("\n")
}
