package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
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

	var ids []string

	input := bufio.NewScanner(infile)
	for input.Scan() {
		ids = append(ids, input.Text())
	}

	var common string

outer:
	for i, s1 := range ids {
		for _, s2 := range ids[i:] {
			match, count := strMatch(s1, s2)
			if count == 1 {
				common = match
				break outer
			}
		}
	}

	fmt.Fprintln(os.Stdout, common)
}

func strMatch(s1, s2 string) (string, int) {
	diffCount := len(s1) - len(s2)

	if diffCount > 0 {
		s1 = s1[:len(s2)]
	}

	if diffCount < 0 {
		s2 = s2[:len(s1)]
		diffCount = -diffCount
	}

	match := make([]rune, 0, len(s1))
	for i, v := range s1 {
		if v == rune(s2[i]) {
			match = append(match, v)
		} else {
			diffCount++
		}
	}

	return string(match), diffCount
}
