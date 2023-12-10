package day_05

import (
	"aoc2023/util"
	"testing"

	"github.com/google/go-cmp/cmp"
)

func Test_parse(t *testing.T) {
	lines := util.GetLinesFromFile("input_small.txt")
	actualSeeds, actualMaps := parse(lines)
	expectedSeeds := Seeds{79, 14, 55, 13}
	expectedMaps := []Map{
		[]Edge{
			{50, 98, 2},
			{52, 50, 48},
		},
		[]Edge{
			{0, 15, 37},
			{37, 52, 2},
			{39, 0, 15},
		},
		[]Edge{
			{49, 53, 8},
			{0, 11, 42},
			{42, 0, 7},
			{57, 7, 4},
		},
		[]Edge{
			{88, 18, 7},
			{18, 25, 70},
		},
		[]Edge{
			{45, 77, 23},
			{81, 45, 19},
			{68, 64, 13},
		},
		[]Edge{
			{0, 69, 1},
			{1, 0, 69},
		},
		[]Edge{
			{60, 56, 37},
			{56, 93, 4},
		},
	}

	diff := cmp.Diff(expectedSeeds, actualSeeds)
	if diff != "" {
		t.Errorf("Seeds\n%v", diff)
	}
	diff = cmp.Diff(expectedMaps, actualMaps)
	if diff != "" {
		t.Errorf("Maps\n%v", diff)
	}
}

func Test_constructSeedBuckets(t *testing.T) {
	lines := util.GetLinesFromFile("input_small.txt")
	seedRanges, _ := parse(lines)
	actual := constructSeedBuckets(seedRanges)
	expected := SeedBuckets{
		[2]int{79, 14},
		[2]int{55, 13},
	}
	diff := cmp.Diff(expected, actual)
	if diff != "" {
		t.Errorf(diff)
	}
}

func Test_findSmallestSeed(t *testing.T) {
	sortedSeedBuckets := SeedBuckets{
		[2]int{10, 5},
		[2]int{16, 5},
	}
	compositeMap := Map{
		[3]int{3, 1, 2},
		[3]int{8, 3, 10},
	}
	actual := findSmallestSeed(sortedSeedBuckets, compositeMap)
	expected := 10
	if actual != expected {
		t.Errorf("Expected %d; Got %d", expected, actual)
	}

	lines := util.GetLinesFromFile("input_small.txt")
	seedRanges, maps := parse(lines)
	seedBuckets := constructSeedBuckets(seedRanges)
	merged := MergeAll(maps)

	actual = findSmallestSeed(seedBuckets, merged)
	expected = 82
	if actual != expected {
		t.Errorf("Expected %d; Got %d", expected, actual)
	}
}

func TestSolve_Part1(t *testing.T) {
	lines := util.GetLinesFromFile("input_small.txt")
	result := Part1(lines)
	expected := 35
	if result != expected {
		t.Errorf("Expected %d; Got %d", expected, result)
	}
}

func TestSolve_Part2(t *testing.T) {
	lines := util.GetLinesFromFile("input_small.txt")
	result := Part2(lines)
	expected := 46
	if result != expected {
		t.Errorf("Expected %d; Got %d", expected, result)
	}
}
