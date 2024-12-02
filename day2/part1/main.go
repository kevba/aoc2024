package main

import (
	"aoc2024"
	"fmt"
	"strings"
)

func main() {
	timer := aoc2024.Time()
	defer timer()

	input := aoc2024.GetInput()
	answer := solve(input)
	fmt.Println(answer)
}

func solve(lines []string) int {
	safeCount := 0

	for _, r := range lines {
		values := aoc2024.AtoiSlice(strings.Split(r, " "))
		direction := 0
		isUnsafe := false

		for i, v := range values[:len(values)-1] {
			diff := v - values[i+1]
			delta := aoc2024.IntAbs(diff)

			if delta <= 0 || delta > 3 {
				isUnsafe = true
				break
			}
			if diff > 0 && direction >= 0 {
				direction = 1
				continue
			}
			if diff < 0 && direction <= 0 {
				direction = -1
				continue
			}
			isUnsafe = true
			break
		}

		if !isUnsafe {
			safeCount += 1
		}

	}

	return safeCount
}
