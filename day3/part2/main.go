package main

import (
	"aoc2024"
	"fmt"
	"regexp"
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
	multipliers := [][]string{}

	multRX := regexp.MustCompile(`mul\(([0-9]+),([0-9]+)\)|do\(\)|don't\(\)`)

	doMode := true
	for _, l := range lines {
		multipliers := append(multipliers, multRX.FindAllStringSubmatch(l, -1)...)
		for _, m := range multipliers {
			if m[0] == "do()" {
				doMode = true
				continue
			}
			if m[0] == "don't()" {
				doMode = false
				continue

			}
			if !doMode {
				continue
			}
			count += (aoc2024.Atoi(m[1]) * aoc2024.Atoi(m[2]))
		}

	}

	return count
}
