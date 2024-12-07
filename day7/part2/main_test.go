package main

import (
	"aoc2024"
	"testing"
)

func TestSolve(t *testing.T) {
	input := aoc2024.GetTestInput()
	solution := 11387

	answer := solve(input)
	if answer != solution {
		t.Errorf("answer %v is not equal to solution %v", answer, solution)
	}
}

func TestOperate(t *testing.T) {
	c := &Calibration{result: 100, values: []int{0, 1, 0, 0, 1}}
	possible := c.operate(1, c.values[0])
	if !possible {
		t.Errorf("expected to be %v", c.result)
	}
}
