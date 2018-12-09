package main

import (
	"strconv"
	"strings"
	"unicode"
)

func intFromString(s string) int {
	i, err := strconv.Atoi(s)
	check(err)
	return i
}

func intFields(s string) []int {
	tokens := strings.FieldsFunc(s, func(c rune) bool {
		return !unicode.IsNumber(c)
	})

	ints := []int{}
	for _, t := range tokens {
		ints = append(ints, intFromString(t))
	}

	return ints
}
