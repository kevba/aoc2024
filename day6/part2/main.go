package main

import (
	"aoc2024"
	"fmt"
	"sync"
)

type Blocker int

const (
	outOfBounds = iota
	freeSpace
	blockedSpace
	looping
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
	runeLines := [][]rune{}
	for _, r := range lines {
		runeLines = append(runeLines, []rune(r))
	}

	d := data{
		startX:         curX,
		startY:         curY,
		startDirection: direction,
	}

	isLoop := make(chan bool)
	count := 0
	var wg sync.WaitGroup

	for y, rows := range runeLines {
		for x, v := range rows {
			if x == d.startX && y == d.startY {
				continue
			}
			if v == '#' {
				continue
			}

			// To improve performance, we could only starts routines for every tile thats actually visited.
			// So do an initial run and only iterate coordinates that have been visited.
			force := func() {
				defer wg.Done()
				newLines := cloneLines(runeLines)
				newLines[y][x] = '0'
				exitCode := d.walk(newLines)
				isLoop <- exitCode == looping
			}
			wg.Add(1)
			go force()
		}
	}

	go func() {
		wg.Wait()
		close(isLoop) // Close the channel when all workers have finished
	}()

	for l := range isLoop {
		if l {
			count += 1
		}
	}

	return count
}

type data struct {
	startX         int
	startY         int
	startDirection int
}

func (d data) walk(lines [][]rune) int {
	curX := d.startX
	curY := d.startY
	direction := d.startDirection

	// Map access is very fast, especially when using integers instead of using strings
	visited := make(map[int]interface{})
	val := curX<<20 | curY<<4 | direction

	visited[val] = nil
	// Assuming that a high count means its looping actually works quite well, a lot quicker then checking the slice :)
	// count := 0

	for {
		// count++
		// if count > 9999{
		// 	return looping
		// }
		nextX := curX + directions[direction][0]
		nextY := curY + directions[direction][1]
		blocked := isBlocked(lines, nextX, nextY)

		if blocked == outOfBounds {
			return outOfBounds
		}

		// Rotate right if the next position is blocking
		if blocked == blockedSpace {
			direction++
			if direction > 3 {
				direction = 0
			}
		}

		// move to the new space if its open
		if blocked == freeSpace {
			curX = nextX
			curY = nextY
		}

		coordsFmt := curX<<20 | curY<<4 | direction
		// If the new conditions are equal to the starting contions we must be looping
		if _, ok := visited[coordsFmt]; ok {
			return looping
		}
		visited[coordsFmt] = nil
	}

}

func isBlocked(in [][]rune, x int, y int) Blocker {
	if y >= 0 && len(in) > y {
		row := in[y]
		if x >= 0 && len(row) > x {
			if row[x] == '#' || row[x] == '0' {
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

func printMap(lines [][]rune) {
	for _, row := range lines {
		for _, r := range row {
			fmt.Printf("%c ", r)
		}
		fmt.Println()
	}
	fmt.Println()

}

func cloneLines(lines [][]rune) [][]rune {
	newLines := make([][]rune, len(lines))
	for y, r := range lines {
		newRows := make([]rune, len(r))
		copy(newRows, r)
		newLines[y] = newRows
	}
	return newLines
}
