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
	count := 0
	rules := [][]int{}
	updates := [][]int{}

	for i, l := range lines {
		if l == "" {
			for _, r := range lines[:i] {
				rn := aoc2024.AtoiSlice(strings.Split(r, "|"))

				rules = append(rules, rn)
			}
			for _, u := range lines[i+1:] {
				un := aoc2024.AtoiSlice(strings.Split(u, ","))
				updates = append(updates, un)
			}
			break
		}
	}

	for _, pages := range updates {
		if !matchRule(rules, pages) {
			pages = sort(rules, pages)
			middle := pages[(len(pages) / 2)]
			count += middle
		}
	}

	return count
}

func matchRule(rules [][]int, pages []int) bool {
	for _, r := range rules {
		if index := slices.Index(pages, r[0]); index > 0 {
			indexBefore := slices.Index(pages[:index], r[1])
			if indexBefore > -1 {
				return false
			}
		}
	}

	return true

}

func sort(rules [][]int, page []int) []int {
	sortF := func(a int, b int) int {
		valuesAfterA := []int{}

		for _, r := range rules {
			if r[0] == a {
				valuesAfterA = append(valuesAfterA, r[1])
			}
		}

		for _, v := range valuesAfterA {
			if v == b {
				return 1
			}
		}
		return -1
	}

	slices.SortFunc(page, sortF)
	return page
}
