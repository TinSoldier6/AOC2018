package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"os"
	"strings"
	"unicode"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Fprintf(os.Stderr, "Usage: %s <inputfile>", os.Args[0])
		os.Exit(1)
	}

	infile, err := ioutil.ReadFile(os.Args[1])
	fail(err)

	polymers := "ABCDEFGHIJKLMNOPQRSTUVWXYZ"
	// Whitespace is a gotcha
	input := strings.TrimSpace(string(infile))
	min := len(input)
	for _, r := range polymers {
		s := remove(r, input)
		output := pass(s)
		if len(output) < min {
			min = len(output)
		}
	}
	fmt.Fprintln(os.Stdout, min)

}

func react(r1, r2 rune) bool {
	return unicode.ToLower(r1) == unicode.ToLower(r2) && unicode.IsLower(r1) != unicode.IsLower(r2)
}

func fail(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

func pass(input string) string {
	in := strings.NewReader(input)
	out := strings.Builder{}

	var reacted bool
	for {
		r1, _, err := in.ReadRune()
		if err == io.EOF {
			break
		}
		fail(err)
		r2, _, err := in.ReadRune()
		if err == io.EOF {
			out.WriteRune(r1)
			break
		}
		fail(err)
		if react(r1, r2) {
			reacted = true
			continue
		}
		_, err = out.WriteRune(r1)
		fail(err)
		err = in.UnreadRune()
		fail(err)
	}

	output := out.String()
	if reacted {
		output = pass(output)
	}

	return output
}

func remove(r1 rune, s string) string {
	r2 := unicode.ToLower(r1)
	out := strings.Builder{}
	for _, r := range s {
		if r == r1 || r == r2 {
			continue
		}
		out.WriteRune(r)
	}
	return out.String()
}
