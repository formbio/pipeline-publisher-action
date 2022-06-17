package main

import (
	"fmt"
	"os"
)

func main() {
	fmt.Println("HELLOOOOOOOOOOOO")

	fmt.Println(os.Getwd())

	files, err := os.ReadDir(".")
	if err != nil {
		panic(err)
	}
	for _, f := range files {
		fmt.Println(f.Info())
	}
}
