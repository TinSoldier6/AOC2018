package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
	"unicode"
)

type rectangle struct {
	id, left, top, width, height int
}

const (
	width = 1000
)

func main() {
	if len(os.Args) < 2 {
		fmt.Fprintf(os.Stderr, "Usage: %s <inputfile>", os.Args[0])
		os.Exit(1)
	}

	infile, err := os.Open(os.Args[1])
	if err != nil {
		log.Fatal(err)
	}

	grid := make(map[int]int)
	input := bufio.NewScanner(infile)
	for input.Scan() {
		nextRect, err := parseRectangle(input.Text())
		if err != nil {
			log.Fatal(err)
		}

		for row := nextRect.top * width; row < (nextRect.top+nextRect.height)*width; row += width {
			for col := nextRect.left; col < nextRect.left+nextRect.width; col++ {
				grid[row+col]++
			}
		}
	}

	var overlap int
	for _, v := range grid {
		if v > 1 {
			overlap++
		}
	}

	fmt.Fprintln(os.Stdout, overlap)

}

func parseRectangle(s string) (rect rectangle, err error) {
	tokens := strings.FieldsFunc(s, func(c rune) bool {
		return !unicode.IsNumber(c)
	})
	if len(tokens) != 5 {
		return rectangle{}, strconv.ErrSyntax
	}

	values := make([]int, 5)
	for i, v := range tokens {
		values[i], err = strconv.Atoi(v)
		if err != nil {
			return rectangle{}, err
		}
	}

	return rectangle{values[0], values[1], values[2], values[3], values[4]}, nil
}
