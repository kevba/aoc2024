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
	freqs := map[rune][][]int{}

	for y, l := range lines {
		for x, f := range l {
			if f == '.' {
				continue
			}
			if _, ok := freqs[f]; !ok {
				freqs[f] = [][]int{}
			}
			freqs[f] = append(freqs[f], []int{x, y})
		}
	}
	height := len(lines) - 1
	width := len(lines[0]) - 1

	antinodes := map[string]interface{}{}

	for _, points := range freqs {
		for _, point := range points {
			for _, point2 := range points {
				if point2[0] == point[0] && point2[1] == point[1] {
					continue
				}

				xD := point2[0] - point[0]
				yD := point2[1] - point[1]

				// A linear functio would have been neat, but this gets the job done just fine
				x := 0
				for {
					antiX := point[0] + xD*-x
					antiY := point[1] + yD*-x
					x++

					if antiX < 0 || antiX > width {
						break
					}
					if antiY < 0 || antiY > height {
						break
					}

					antinodes[fmt.Sprintf("%v,%v", antiX, antiY)] = nil
				}
			}
		}
	}

	return len(antinodes)
}
