package main

import "testing"

// TestPart1 with example input
func TestPart1(t *testing.T) {
	part1In := []string{}
	part1Expect := "0"

	if got := part1(part1In); got != part1Expect {
		t.Errorf("got: %s, wanted: %s\n", got, part1Expect)
	}
}

// TestPart2 with example input
func TestPart2(t *testing.T) {
	part2In := []string{}
	part2Expect := "0"

	if got := part1(part2In); got != part2Expect {
		t.Errorf("got: %s, wanted: %s\n", got, part2Expect)
	}
}
