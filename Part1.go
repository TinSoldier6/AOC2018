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

	matchTwos := 0
	matchThrees := 0
	input := bufio.NewScanner(infile)
	for input.Scan() {
		s := input.Text()
		if hasMatches(s, 2) {
			matchTwos++
		}
		if hasMatches(s, 3) {
			matchThrees++
		}
	}

	fmt.Fprint(os.Stdout, matchTwos*matchThrees)
}

// hasMatches returns true if s has any groups of exactly c of any character
func hasMatches(s string, c int) bool {
	m := make(map[rune]int)
	r := []rune(s)

	for _, v := range r {
		m[v]++
	}

	for _, v := range m {
		if v == c {
			return true
		}
	}

	return false
}
