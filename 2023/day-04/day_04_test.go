package main

import (
	"aoc2023/util"
	"testing"

	"github.com/google/go-cmp/cmp"
)

func Test_parse(t *testing.T) {
	lines := util.GetLinesFromFile("input_small.txt")
	actual := parse(lines)
	expected := []Card{
		{
			WinningNumbers: []int{41, 48, 83, 86, 17},
			Numbers:        map[int]struct{}{83: {}, 86: {}, 6: {}, 31: {}, 17: {}, 9: {}, 48: {}, 53: {}},
		},
		{
			WinningNumbers: []int{13, 32, 20, 16, 61},
			Numbers:        map[int]struct{}{61: {}, 30: {}, 68: {}, 82: {}, 17: {}, 32: {}, 24: {}, 19: {}},
		},
		{
			WinningNumbers: []int{1, 21, 53, 59, 44},
			Numbers:        map[int]struct{}{69: {}, 82: {}, 63: {}, 72: {}, 16: {}, 21: {}, 14: {}, 1: {}},
		},
		{
			WinningNumbers: []int{41, 92, 73, 84, 69},
			Numbers:        map[int]struct{}{59: {}, 84: {}, 76: {}, 51: {}, 58: {}, 5: {}, 54: {}, 83: {}},
		},
		{
			WinningNumbers: []int{87, 83, 26, 28, 32},
			Numbers:        map[int]struct{}{88: {}, 30: {}, 70: {}, 12: {}, 93: {}, 22: {}, 82: {}, 36: {}},
		},
		{
			WinningNumbers: []int{31, 18, 13, 56, 72},
			Numbers:        map[int]struct{}{74: {}, 77: {}, 10: {}, 23: {}, 35: {}, 67: {}, 36: {}, 11: {}},
		},
	}

	diff := cmp.Diff(expected, actual)
	if diff != "" {
		t.Error(diff)
	}
}

func TestSolve_Part1(t *testing.T) {
	lines := util.GetLinesFromFile("input_small.txt")
	result := part1(lines)
	expected := 13
	if result != expected {
		t.Errorf("Expected %d; Got %d", expected, result)
	}
}

func TestSolve_Part2(t *testing.T) {
	lines := util.GetLinesFromFile("input_small.txt")
	result := part2(lines)
	expected := 30
	if result != expected {
		t.Errorf("Expected %d; Got %d", expected, result)
	}
}
