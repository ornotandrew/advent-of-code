package day_09

import (
	"aoc2023/util"
	"testing"

	"github.com/google/go-cmp/cmp"
)

func TestSolve_Part1(t *testing.T) {
	lines := util.GetLinesFromFile("input_small.txt")
	result := Part1(lines)
	expected := 114
	if result != expected {
		t.Errorf("Expected %d; Got %d", expected, result)
	}
}

func TestSolve_Part2(t *testing.T) {
	lines := util.GetLinesFromFile("input_small.txt")
	result := Part2(lines)
	expected := 2
	if result != expected {
		t.Errorf("Expected %d; Got %d", expected, result)
	}
}

func Test_parse(t *testing.T) {
	lines := util.GetLinesFromFile("input_small.txt")
	actual := parse(lines)
	expected := []Sequence{
		{0, 3, 6, 9, 12, 15},
		{1, 3, 6, 10, 15, 21},
		{10, 13, 16, 21, 30, 45},
	}
	diff := cmp.Diff(expected, actual)
	if diff != "" {
		t.Error(diff)
	}
}
