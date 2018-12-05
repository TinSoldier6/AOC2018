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

func (r1 rectangle) union(r2 rectangle) rectangle {
	if r1.left > r2.left+r2.width || r2.left > r1.left+r1.width {
		return rectangle{}
	}
	if r1.top > r2.top+r2.height || r2.top > r1.top+r1.height {
		return rectangle{}
	}
	left := max(r1.left, r2.left)
	top := max(r1.top, r2.top)
	right := min(r1.left+r1.width, r2.left+r2.width)
	bottom := min(r1.top+r1.height, r2.top+r2.height)
	return rectangle{0, left, top, right - left, bottom - top}
}

func (r1 rectangle) area() int {
	return r1.width * r1.height
}

func main() {
	if len(os.Args) < 2 {
		fmt.Fprintf(os.Stderr, "Usage: %s <inputfile>", os.Args[0])
		os.Exit(1)
	}

	infile, err := os.Open(os.Args[1])
	if err != nil {
		log.Fatal(err)
	}

	rectangles := make([]rectangle, 0)
	input := bufio.NewScanner(infile)
	for input.Scan() {
		nextRect, err := parseRectangle(input.Text())
		if err != nil {
			log.Fatal(err)
		}
		rectangles = append(rectangles, nextRect)
	}

	var overlap int
	for i, r1 := range rectangles {
		for _, r2 := range rectangles[i:] {
			area := r1.union(r2).area()
			fmt.Printf("%d ", area)
			overlap += area
		}
	}

	fmt.Fprintln(os.Stdout, overlap)

}

func max(a, b int) int {
	if a >= b {
		return a
	}
	return b
}

func min(a, b int) int {
	if a <= b {
		return a
	}
	return b
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
