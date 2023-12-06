package main

import (
	"aoc2023/util"
	"testing"

	"github.com/google/go-cmp/cmp"
)

func Test_findNumbers(t *testing.T) {
	lines := util.GetLinesFromFile("input_1_small.txt")
	actual := findNumbers(lines)
	expected := []Number{
		{0, 0, 2, 467},
		{0, 5, 7, 114},
		{2, 2, 3, 35},
		{2, 6, 8, 633},
		{4, 0, 2, 617},
		{5, 7, 8, 58},
		{6, 2, 4, 592},
		{7, 6, 8, 755},
		{9, 1, 3, 664},
		{9, 5, 7, 598},
	}

	diff := cmp.Diff(expected, actual)
	if diff != "" {
		t.Errorf(diff)
	}
}

func Test_findSymbols(t *testing.T) {
	lines := util.GetLinesFromFile("input_1_small.txt")
	actual := findSymbols(lines)
	expected := []Symbol{
		{1, 3, '*'},
		{3, 6, '#'},
		{4, 3, '*'},
		{5, 5, '+'},
		{8, 3, '$'},
		{8, 5, '*'},
	}

	diff := cmp.Diff(expected, actual)
	if diff != "" {
		t.Errorf(diff)
	}
}

func Test_numberIsAdjacentToSymbol(t *testing.T) {
	lines := util.GetLinesFromFile("input_1_small.txt")
	numbers := findNumbers(lines)
	symbols := findSymbols(lines)

	type Test struct {
		Number   Number
		Expected bool
	}
	tests := []Test{
		{numbers[0], true},
		{numbers[1], false},
		{numbers[2], true},
		{numbers[3], true},
		{numbers[4], true},
		{numbers[5], false},
		{numbers[6], true},
		{numbers[7], true},
		{numbers[8], true},
		{numbers[9], true},
	}

	for _, test := range tests {
		if numberIsAdjacentToSymbol(test.Number, symbols) != test.Expected {
			t.Errorf("Expected %d to be %v", test.Number, test.Expected)
		}
	}

}

func Test_getAdjacentNumbers(t *testing.T) {
	lines := util.GetLinesFromFile("input_1_small.txt")
	numbers := findNumbers(lines)
	symbols := findSymbols(lines)

	type Test struct {
		Symbol   Symbol
		Expected int
	}

	tests := []Test{
		{symbols[0], 2},
		{symbols[2], 1},
		{symbols[5], 2},
	}

	for _, test := range tests {
		actual := len(getAdjacentNumbers(test.Symbol, numbers))
		if actual != test.Expected {
			t.Errorf("Expected %d to result in %v but got %v", test.Symbol, test.Expected, actual)
		}
	}

}

func TestSolve_Part1(t *testing.T) {
	lines := util.GetLinesFromFile("input_1_small.txt")
	result := part1(lines)
	expected := 4361
	if result != expected {
		t.Errorf("Expected %d; Got %d", expected, result)
	}
}

func TestSolve_Part2(t *testing.T) {
	lines := util.GetLinesFromFile("input_1_small.txt")
	result := part2(lines)
	expected := 467835
	if result != expected {
		t.Errorf("Expected %d; Got %d", expected, result)
	}
}
