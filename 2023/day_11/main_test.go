package day_11

import (
	"aoc2023/util"
	"testing"

	"github.com/google/go-cmp/cmp"
)

func TestSolve_Part1(t *testing.T) {
	lines := util.GetLinesFromFile("input_small.txt")
	result := Part1(lines)
	expected := 374
	if result != expected {
		t.Errorf("Expected %d; Got %d", expected, result)
	}
}

func TestSolve_Part2(t *testing.T) {
	lines := util.GetLinesFromFile("input_small.txt")
	result := Part2(lines)
	expected := 82000210
	if result != expected {
		t.Errorf("Expected %d; Got %d", expected, result)
	}
}

func Test_parse(t *testing.T) {
	lines := util.GetLinesFromFile("input_small.txt")
	actual := parse(lines)
	expected := Image{
		map[Pos]struct{}{
			{0, 3}: {},
			{1, 7}: {},
			{2, 0}: {},
			{4, 6}: {},
			{5, 1}: {},
			{6, 9}: {},
			{8, 7}: {},
			{9, 0}: {},
			{9, 4}: {},
		},
		10,
		10,
	}
	if diff := cmp.Diff(expected, actual); diff != "" {
		t.Error(diff)
	}
}

func Test_Expand(t *testing.T) {
	image := parse(util.GetLinesFromFile("input_small.txt"))
	actual := image.Expand(2)
	expected := Image{
		map[Pos]struct{}{
			{0, 4}:  {},
			{1, 9}:  {},
			{2, 0}:  {},
			{5, 8}:  {},
			{6, 1}:  {},
			{7, 12}: {},
			{10, 9}: {},
			{11, 0}: {},
			{11, 5}: {},
		},
		0,
		0,
	}

	if diff := cmp.Diff(expected, actual); diff != "" {
		t.Error(diff)
	}
}
