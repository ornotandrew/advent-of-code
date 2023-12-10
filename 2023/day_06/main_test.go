package day_06

import (
	"aoc2023/util"
	"testing"

	"github.com/google/go-cmp/cmp"
)

func Test_parse_1(t *testing.T) {
	lines := util.GetLinesFromFile("input_small.txt")
	actual := parse_1(lines)
	expected := []Race{
		{7, 9},
		{15, 40},
		{30, 200},
	}
	diff := cmp.Diff(expected, actual)
	if diff != "" {
		t.Error(diff)
	}
}

func Test_parse_2(t *testing.T) {
	lines := util.GetLinesFromFile("input_small.txt")
	actual := parse_2(lines)
	expected := Race{71530, 940200}
	diff := cmp.Diff(expected, actual)
	if diff != "" {
		t.Error(diff)
	}
}

func Test_roots(t *testing.T) {
	lines := util.GetLinesFromFile("input_small.txt")
	races := parse_1(lines)
	expected := [][2]int{
		{2, 5},
		{4, 11},
		{11, 19},
	}
	for i, race := range races {
		minChargeTime, maxChargeTime := roots(race)
		actual := [2]int{minChargeTime, maxChargeTime}
		if actual != expected[i] {
			t.Errorf("Race %v; Expected %v; Got %v\n", i, expected[i], actual)
		}
	}
}

func TestSolve_Part1(t *testing.T) {
	lines := util.GetLinesFromFile("input_small.txt")
	result := Part1(lines)
	expected := 288
	if result != expected {
		t.Errorf("Expected %d; Got %d", expected, result)
	}
}

func TestSolve_Part2(t *testing.T) {
	lines := util.GetLinesFromFile("input_small.txt")
	result := Part2(lines)
	expected := 71503
	if result != expected {
		t.Errorf("Expected %d; Got %d", expected, result)
	}
}
