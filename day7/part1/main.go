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
	count := 0

	for _, l := range lines {
		cal := parseCalibration(l)
		if cal.possiblyTrue() {
			count += cal.result
		}
	}

	return count

}

func parseCalibration(l string) *Calibration {
	parts := strings.Split(l, " ")

	return &Calibration{
		result: aoc2024.Atoi(parts[0][:len(parts[0])-1]),
		values: aoc2024.AtoiSlice(parts[1:]),
	}
}

type Calibration struct {
	result int
	values []int
}

func (c *Calibration) possiblyTrue() bool {
	isSolvable := c.operate(1, c.values[0])
	return isSolvable
}

func (c *Calibration) operate(index int, total int) bool {
	if index >= len(c.values) {
		return total == c.result
	}

	multiplied := total * c.values[index]
	summed := total + c.values[index]
	if c.operate(index+1, summed) {
		return true
	}
	if c.operate(index+1, multiplied) {
		return true
	}
	return false
}
