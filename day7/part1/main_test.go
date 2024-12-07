package main

import (
	"aoc2024"
	"testing"
)

func TestSolve(t *testing.T) {
	input := aoc2024.GetTestInput()
	solution := 3749

	answer := solve(input)
	if answer != solution {
		t.Errorf("answer %v is not equal to solution %v", answer, solution)
	}
}

func TestOperate(t *testing.T) {
	c := &Calibration{result: 100, values: []int{24, 2, 50, 2, 0, 0}}
	possible := c.operate(1, c.values[0])
	if !possible {
		t.Errorf("expected to be %v", c.result)
	}
}
