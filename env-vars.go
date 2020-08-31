package main

import (
	"fmt"
	"os"
)

func main() {

	for _, env := range os.Environ() {
		fmt.Println(env)
	}

	name := os.Getenv("USERNAME")
	course := "Docker Deep Dive"

	fmt.Println("\nHi", name, "you're currently watching", course)

}
