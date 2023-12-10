package day_01

import (
	"aoc2023/util"
	"testing"
)

func TestSolve_Part1(t *testing.T) {
	lines := util.GetLinesFromFile("input_1_small.txt")

	answer := Part1(lines)
	if answer != 142 {
		t.Errorf("Expected 142; Got %d", answer)
	}

	answer = Part1([]string{"7ggzdnjxndfive"})
	if answer != 77 {
		t.Errorf("Expected 77; Got %d", answer)
	}
}

func TestSolve_Part2(t *testing.T) {
	lines := util.GetLinesFromFile("input_2_small.txt")
	answer := Part2(lines)
	expected := 281
	if answer != expected {
		t.Errorf("Expected %d; Got %d", expected, answer)
	}
}
