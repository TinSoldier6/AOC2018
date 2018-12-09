// Template for each day's code
package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
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
func part1(input []string) (output string) {
	x, y := []int{}, []int{}
	for _, s := range input {
		coords := strings.Split(s, ", ")
		xs, err := strconv.Atoi(coords[0])
		if err != nil {
			log.Fatal(err)
		}
		ys, err := strconv.Atoi(coords[1])
		if err != nil {
			log.Fatal(err)
		}
		x, y = append(x, xs), append(y, ys)
	}

	xmin, ymin := min(x...), min(y...)
	xmax, ymax := max(x...), max(y...)
	width := xmax - xmin
	height := ymax - ymin

	grid := make([][]int, height)
	for i := 0; i < height; i++ {
		grid[i] = make([]int, width)
	}

	return output
}

// part2 takes a list of strings and computes the answer for part 2 of the problem.
func part2(input []string) (output string) {
	return output
}

func abs(a int) int {
	b := a >> (strconv.IntSize - 1)
	return (a ^ b) - b
}

const maxint = (1<<(strconv.IntSize-1) - 1)

func min(a ...int) (b int) {
	b = maxint
	for _, v := range a {
		if v < b {
			b = v
		}
	}
	return b
}

func max(a ...int) (b int) {
	for _, v := range a {
		if v > b {
			b = v
		}
	}
	return b
}

func distance(x1, y1, x2, y2 int) int {
	return abs(x2-x1) + abs(y2-y1)
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
