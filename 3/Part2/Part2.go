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
	overlapped                   bool
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

	rectangles := make([]rectangle, 0)
	input := bufio.NewScanner(infile)
	for input.Scan() {
		nextRect, err := parseRectangle(input.Text())
		if err != nil {
			log.Fatal(err)
		}
		rectangles = append(rectangles, nextRect)
	}

	for i := 0; i < len(rectangles); i++ {
		for j := i + 1; j < len(rectangles); j++ {
			if overlaps(rectangles[i], rectangles[j]) {
				rectangles[i].overlapped = true
				rectangles[j].overlapped = true
			}
		}
	}

	var id int
	for _, r := range rectangles {
		if !r.overlapped {
			id = r.id
		}
	}

	fmt.Fprintln(os.Stdout, id)

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

	return rectangle{values[0], values[1], values[2], values[3], values[4], false}, nil
}

func overlaps(r1, r2 rectangle) bool {
	if r1.left > r2.left+r2.width || r2.left > r1.left+r1.width {
		return false
	}
	if r1.top > r2.top+r2.height || r2.top > r1.top+r1.height {
		return false
	}
	return true
}
