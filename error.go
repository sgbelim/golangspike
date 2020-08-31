package main

import (
	"fmt"
	"os"
)

func main() {
	_, err := os.Open("c:\\temp\\test.txt")

	if err != nil {
		fmt.Println("Error returnd was:", err)
	}
}
