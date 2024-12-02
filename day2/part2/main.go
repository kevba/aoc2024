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

		if isValid(values) < 0 {
			safeCount += 1
			continue
		}

		for i := range values {
			pos1 := aoc2024.RemoveIndex(append([]int{}, values...), i)
			if isValid(pos1) < 0 {
				safeCount += 1
				break
			}
		}
	}

	return safeCount
}

func isValid(values []int) int {
	direction := 0
	isUnsafe := -1

	for i, v := range values[:len(values)-1] {
		diff := v - values[i+1]
		delta := aoc2024.IntAbs(diff)

		if delta > 0 && delta < 4 {
			if diff > 0 && direction >= 0 {
				direction = 1
				continue
			}
			if diff < 0 && direction <= 0 {
				direction = -1
				continue
			}
		}

		isUnsafe = i
		break
	}

	return isUnsafe
}
