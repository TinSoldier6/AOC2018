// Template for each day's code
package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

// read reads in a file and returns each line as an element of []string.
func read(file string) (lines []string, err error) {
	f, err := os.Open(file)
	if err != nil {
		return
	}
	defer f.Close()

	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}
	err = scanner.Err()
	return
}

// part1 takes a list of strings and computes the answer for part 1 of the problem.
func part1(input []string) string {

}

// part2 takes a list of strings and computes the answer for part 2 of the problem.
func part2(input []string) string {

}

func main() {
	if len(os.Args) < 2 {
		fmt.Fprintf(os.Stderr, "Usage: %s <inputfile>", os.Args[0])
		os.Exit(1)
	}

	input, err := read(os.Args[1])
	if err != nil {
		log.Fatal(err)
	}
	fmt.Fprintf(os.Stdout, "Part 1\n%s\n\n", part1(input))
	fmt.Fprintf(os.Stdout, "Part 2\n%s\n\n", part2(input))
}
