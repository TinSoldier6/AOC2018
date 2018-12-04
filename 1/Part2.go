package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
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

	var list []int
	var seen = make(map[int]bool)
	var finished bool
	seen[0] = true

	input := bufio.NewScanner(infile)
	frequency := 0
	for input.Scan() {
		change, err := strconv.Atoi(input.Text())
		if err != nil {
			log.Fatal(err)
		}
		frequency += change
		list = append(list, change)
		if seen[frequency] {
			finished = true
		}
		seen[frequency] = true
	}

	for !finished {
		for _, v := range list {
			frequency += v
			if seen[frequency] {
				finished = true
				break
			}
			seen[frequency] = true
		}
	}

	fmt.Fprintln(os.Stdout, frequency)
}
