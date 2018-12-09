package main

import (
	"fmt"
	"os"
)

// This is the main driver of the program.
func main() {
	if len(os.Args) < 2 {
		fmt.Fprintf(os.Stderr, "Usage: %s <inputfile>", os.Args[0])
		os.Exit(1)
	}

	input := read(os.Args[1])

	fmt.Fprintf(os.Stdout, "Part 1:\n%s\n\n", part1(input))
	fmt.Fprintf(os.Stdout, "Part 2:\n%s\n\n", part2(input))
}
