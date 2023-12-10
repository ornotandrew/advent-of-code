package day_02

import (
	"aoc2023/util"
	"testing"

	"github.com/google/go-cmp/cmp"
)

func Test_Parse(t *testing.T) {
	lines := util.GetLinesFromFile("input_small.txt")
	result := parse(lines)
	expected := []Game{
		//   r  g  b
		{
			{4, 0, 3},
			{1, 2, 6},
			{0, 2, 0},
		},
		{
			{0, 2, 1},
			{1, 3, 4},
			{0, 1, 1},
		},
		{
			{20, 8, 6},
			{4, 13, 5},
			{1, 5, 0},
		},
		{
			{3, 1, 6},
			{6, 3, 0},
			{14, 3, 15},
		},
		{
			{6, 3, 1},
			{1, 2, 2},
		},
	}

	diff := cmp.Diff(result, expected)
	if diff != "" {
		t.Errorf("Did not match.\n%s", diff)
	}
}

func TestSolve_Part1(t *testing.T) {
	lines := util.GetLinesFromFile("input_small.txt")
	answer := Part1(lines)
	expected := 8
	if answer != expected {
		t.Errorf("Expected %d; Got %d", expected, answer)
	}
}

func TestSolve_Part2(t *testing.T) {
	lines := util.GetLinesFromFile("input_small.txt")
	answer := Part2(lines)
	expected := int64(2286)
	if answer != expected {
		t.Errorf("Expected %d; Got %d", expected, answer)
	}
}
