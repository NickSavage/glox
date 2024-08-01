package main

import (
	"fmt"
	"log"
	"os"
)

// complete main by taking in command line arguments and returning
func main() {
	if len(os.Args) < 2 {
		log.Fatal("You need to pass an argument")
	}

	// You can then handle the arguments as you see fit.
	for _, arg := range os.Args[1:] {
		fmt.Println(arg)
	}
}
