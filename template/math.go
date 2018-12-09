package main

import "strconv"

const maxInt = (1 << (strconv.IntSize - 1)) - 1

func intabs(a int) int {
	b := a >> (strconv.IntSize - 1)
	return (a ^ b) - b
}

func intMin(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func intMax(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func sliceMin(a []int) (b int) {
	b = maxInt
	for _, v := range a {
		if v < b {
			b = v
		}
	}
	return b
}

func sliceMax(a []int) (b int) {
	for _, v := range a {
		if v > b {
			b = v
		}
	}
	return b
}
