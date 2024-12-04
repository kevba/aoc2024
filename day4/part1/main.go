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
			if c == rune('X') {
				count += findXmas(lines, x, y)
			}
		}
	}

	return count
}

func findXmas(lines []string, x int, y int) int {
	directions := [][]int{}

	// find the m to get the right direction
	for _, pattern := range match {
		next := getToken(lines, x+pattern[0], y+pattern[1])
		if next == 'M' {
			directions = append(directions, pattern)
		}
	}

	if len(directions) == 0 {
		return 0
	}

	count := 0
	for _, dir := range directions {
		isComplete := true
		for i, c := range []rune{'A', 'S'} {
			next := getToken(lines, x+dir[0]*(i+2), y+dir[1]*(i+2))
			if next != c {
				isComplete = false
				break
			}
		}

		if isComplete {
			count++
		}
	}

	return count

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

var match = [][]int{
	{-1, -1}, // backwards up
	{-1, 0},  // backwards
	{-1, 1},  // backwards down
	{0, 1},   // down
	{0, -1},  // up
	{1, 1},   // forwards down
	{1, 0},   // forwards
	{1, -1},  // forwards up
}
