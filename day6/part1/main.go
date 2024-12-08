package main

import (
	"aoc2024"
	"fmt"
	"slices"
)

type Blocker int

const (
	outOfBounds = iota
	freeSpace
	blockedSpace
)

func main() {
	timer := aoc2024.Time()
	defer timer()

	input := aoc2024.GetInput()
	answer := solve(input)
	fmt.Println(answer)
}

func solve(lines []string) int {
	curX := 0
	curY := 0
	direction := 0

	for y, row := range lines {
		for x, c := range row {
			if c == rune('^') {
				curX = x
				curY = y
			}
		}
	}

	visited := []string{fmt.Sprintf("%v,%v", curX, curY)}

	for {
		nextX := curX + directions[direction][0]
		nextY := curY + directions[direction][1]
		blocked := isBlocked(lines, nextX, nextY)

		if blocked == outOfBounds {
			break
		}

		if blocked == freeSpace {
			curX = nextX
			curY = nextY
			coordsFmt := fmt.Sprintf("%v,%v", curX, curY)
			if slices.Index(visited, coordsFmt) == -1 {
				visited = append(visited, coordsFmt)
			}
			continue
		}

		if blocked == blockedSpace {
			direction++
			if direction > 3 {
				direction = 0
			}
		}
	}

	return len(visited)
}

func isBlocked(in []string, x int, y int) Blocker {
	if y >= 0 && len(in) > y {
		row := in[y]
		if x >= 0 && len(row) > x {
			if row[x] == '#' {
				return blockedSpace
			}
			return freeSpace
		}
	}

	return outOfBounds
}

var directions = [][]int{
	{0, -1}, // up
	{1, 0},  // forwards
	{0, 1},  // down
	{-1, 0}, // backwards
}

func printMap(lines []string, visited []string) {
	out := [][]rune{}
	for _, r := range lines {
		out = append(out, []rune(r))
	}

	for y, row := range out {
		for x := range row {
			coordsFmt := fmt.Sprintf("%v,%v", x, y)
			if slices.Index(visited, coordsFmt) > -1 {
				out[y][x] = 'x'
			}
		}
	}

	for _, row := range out {
		for _, r := range row {
			fmt.Printf("%c ", r)
		}
		fmt.Println()
	}
	fmt.Println()

}
