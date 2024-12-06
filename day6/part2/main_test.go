package main

import (
	"aoc2024"
	"testing"
)

func TestSolve(t *testing.T) {
	input := aoc2024.GetTestInput()
	solution := 6

	answer := solve(input)
	if answer != solution {
		t.Errorf("answer %v is not equal to solution %v", answer, solution)
	}
}
