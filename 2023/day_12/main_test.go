package day_12

import (
	"aoc2023/util"
	"testing"

	"github.com/google/go-cmp/cmp"
)

func TestSolve_Part1(t *testing.T) {
	lines := util.GetLinesFromFile("input_small.txt")
	result := Part1(lines)
	expected := 21
	if result != expected {
		t.Errorf("Expected %d; Got %d", expected, result)
	}

	lines = util.GetLinesFromFile("input.txt")
	result = Part1(lines)
	if result >= 18388 {
		t.Error("Result too high")
	}
}

func TestSolve_Part2(t *testing.T) {
	lines := util.GetLinesFromFile("input_small.txt")
	result := Part2(lines)
	expected := 525152
	if result != expected {
		t.Errorf("Expected %d; Got %d", expected, result)
	}
}

func Test_FindSolutions(t *testing.T) {
	type Test struct {
		row      Row
		expected int
	}

	for i, test := range []Test{
		{Row{"???.###", []int{1, 1, 3}}, len([]string{"#.#.###"})},
		{Row{".??..??...?##.", []int{1, 1, 3}}, len([]string{
			".#...#....###.",
			".#....#...###.",
			"..#..#....###.",
			"..#...#...###.",
		})},
		{Row{"?#?#?#?#?#?#?#?", []int{1, 3, 1, 6}}, len([]string{
			".#.###.#.######",
		})},
		{Row{"?###????????", []int{3, 2, 1}}, len([]string{
			".###.##.#...",
			".###.##..#..",
			".###.##...#.",
			".###.##....#",
			".###..##.#..",
			".###..##..#.",
			".###..##...#",
			".###...##.#.",
			".###...##..#",
			".###....##.#",
		})},
	} {
		actual := solve(&State{}, test.row, 0, 0, 0)
		if test.expected != actual {
			t.Errorf("Case %d: Expected: %v; Actual: %v\n", i, test.expected, actual)
		}
	}
}

func Test_Expand(t *testing.T) {
	type Test struct {
		row      Row
		expected Row
	}

	for _, test := range []Test{
		{Row{".#", []int{1}}, Row{".#?.#?.#?.#?.#", []int{1, 1, 1, 1, 1}}},
		{Row{"???.###", []int{1, 1, 3}}, Row{"???.###????.###????.###????.###????.###", []int{1, 1, 3, 1, 1, 3, 1, 1, 3, 1, 1, 3, 1, 1, 3}}},
	} {
		actual := test.row.Expand()
		if diff := cmp.Diff(test.expected, actual); diff != "" {
			t.Errorf(diff)
		}
	}
}
