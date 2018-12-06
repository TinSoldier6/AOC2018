package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"sort"
	"strconv"
	"strings"
	"time"
)

// Get input
// Sort input by date and time
// Find how much time each guard slept
// Find which minute guard slept most
// Multiply ID by minute -- output

type event struct {
	time        time.Time
	description string
}

type byTime []event

func (e byTime) Len() int {
	return len(e)
}

func (e byTime) Less(i, j int) bool {
	return e[i].time.Before(e[j].time)
}

func (e byTime) Swap(i, j int) {
	e[i], e[j] = e[j], e[i]
}

type guard struct {
	id, asleep int
	minutes    [60]int
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

	// This program relies on well-formed data

	// Parse the events in the inputs
	events := make([]event, 0)
	tformat := "[2006-01-02 15:04]"
	input := bufio.NewScanner(infile)
	for input.Scan() {
		str := input.Text()
		t, err := time.Parse(tformat, str[:18])
		if err != nil {
			log.Fatal(err)
		}
		events = append(events, event{t, str[19:]})
	}

	sort.Sort(byTime(events))

	// Find out how much each guard slept, and which minutes they were asleep
	var id, asleep, awake int
	guards := make(map[int]guard)
	for _, e := range events {
		str := strings.Fields(e.description)
		if str[0] == "Guard" {
			id, err = strconv.Atoi(str[1][1:])
			if err != nil {
				log.Fatal(err)
			}
		}
		if str[0] == "falls" {
			asleep = e.time.Minute()
		}
		if str[0] == "wakes" {
			awake = e.time.Minute()
			g := guards[id]
			g.id = id
			g.asleep += awake - asleep
			for i := asleep; i < awake; i++ {
				g.minutes[i]++
			}
			guards[id] = g
		}
	}

	// Find out which guard slept the most
	var candidate guard
	for _, g := range guards {
		if g.asleep > candidate.asleep {
			candidate = g
		}
	}

	// Then find which minute has the greatest frequency
	var max, minute int
	for i, m := range candidate.minutes[:] {
		if m > max {
			max = m
			minute = i
		}
	}

	fmt.Fprintln(os.Stdout, minute*candidate.id)
}
