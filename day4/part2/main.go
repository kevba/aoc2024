package main

import (
	"aoc2024"
	"fmt"
)

func main() {
	timer := aoc2024.Time()
	defer timer()

	input := aoc2024.GetInput()
	answer := solve(input)
	fmt.Println(answer)
}

func solve(lines []string) int {
	count := 0

	for y, row := range lines {
		for x, c := range row {
			if c == rune('A') {
				count += findXmas(lines, x, y)
			}
		}
	}

	return count
}

var matchPatterns = [][]int{
	{-1, -1}, // backwards up
	{-1, 1},  // backwards down
}

func findXmas(lines []string, x int, y int) int {
	for _, pattern := range matchPatterns {
		nX := x + pattern[0]
		nY := y + pattern[1]
		pX := x + pattern[0]*-1
		pY := y + pattern[1]*-1
		next := getToken(lines, nX, nY)

		if next == 'M' && getToken(lines, pX, pY) == 'S' {
			continue
		} else if next == 'S' && getToken(lines, pX, pY) == 'M' {
			continue
		} else {
			return 0
		}

	}

	return 1
}

func getToken(in []string, x int, y int) rune {
	if y >= 0 && len(in) > y {

		row := in[y]
		if x >= 0 && len(row) > x {
			return rune(row[x])
		}
	}

	return '.'
}
