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

	input := bufio.NewScanner(infile)
	frequency := 0
	for input.Scan() {
		change, err := strconv.Atoi(input.Text())
		if err != nil {
			log.Fatal(err)
		}
		frequency += change
	}

	fmt.Fprint(os.Stdout, frequency)

}
