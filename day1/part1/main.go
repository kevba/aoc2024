package main

import (
	"aoc2024"
	"fmt"
	"slices"
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
	sumVal := 0
	first := make([]int, len(lines))
	second := make([]int, len(lines))

	for i, l := range lines {
		splitted := strings.Split(l, "   ")

		first[i] = aoc2024.Atoi(splitted[0])
		second[i] = aoc2024.Atoi(splitted[1])
	}

	slices.Sort(first)
	slices.Sort(second)

	for i, v := range first {
		diff := aoc2024.IntAbs(v - second[i])
		sumVal += diff
	}

	return sumVal
}
