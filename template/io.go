package main

import (
	"bufio"
	"log"
	"os"
)

// read reads in a file and returns each line as an element of []string.
func read(file string) (lines []string) {
	f, err := os.Open(file)
	check(err)
	defer f.Close()

	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}
	err = scanner.Err()
	return
}

func check(err error) {
	if err != nil {
		log.Fatal(err)
	}
}
