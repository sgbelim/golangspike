package main

import (
	"fmt"
)

func main() {
	firstRank := "39"
	secondRank := "614"

	if firstRank < secondRank {
		fmt.Println("\nFirst course is doing better than second course")
	} else if firstRank > secondRank {
		fmt.Println("\nOh dear..youe First course must be doing abysmally!")
	} else {
		fmt.Println("\nBoth courses are either the same or something weird is going on")
	}

	if firstRank, secondRank := 39, 614; firstRank < secondRank {
		fmt.Println("\nFirst course is doing better than second course")
	} else if firstRank > secondRank {
		fmt.Println("\nOh dear..youe First course must be doing abysmally!")
	} else {
		fmt.Println("\nBoth courses are either the same or something weird is going on")
	}

}
