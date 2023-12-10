package day_05

import (
	"aoc2023/util"
	"testing"

	"github.com/google/go-cmp/cmp"
)

func Test_MaxValueInSameSrcRange(t *testing.T) {
	type Test struct {
		mapDefinition Map
		input         int
		expected      int
	}

	// https://link.excalidraw.com/l/5SN39rbzyFm/4eTt0xgQYte
	a := Map{Edge{1, 2, 3}}
	b := Map{Edge{4, 2, 3}}
	for _, test := range []Test{
		{a, 1, 1},
		{a, 2, 4},
		{a, 3, 4},
		{a, 4, 4},
		{a, 5, 6},
		{a, 6, 6},

		{b, 1, 1},
		{b, 2, 4},
		{b, 3, 4},
		{b, 4, 4},
		{b, 5, 6},
		{b, 6, 6},
	} {
		actual := test.mapDefinition.MaxValueInSameSrcRange(test.input, 6)
		if actual != test.expected {
			t.Errorf("given: %v; expected: %v; actual: %v\n", test.input, test.expected, actual)
		}
	}

	lines := util.GetLinesFromFile("input_small.txt")
	_, maps := parse(lines)
	seedToSoil := maps[0]
	actual := seedToSoil.MaxValueInSameSrcRange(98, 99)
	expected := 99
	diff := cmp.Diff(expected, actual)
	if diff != "" {
		t.Errorf(diff)
	}

}

func Test_EdgeContaining(t *testing.T) {
	lines := util.GetLinesFromFile("input_small.txt")
	_, maps := parse(lines)
	soilToFertilizer := maps[1]
	actual := soilToFertilizer.EdgeContaining(15)
	expected := Edge{0, 15, 37}
	diff := cmp.Diff(expected, actual)
	if diff != "" {
		t.Errorf(diff)
	}
}

func Test_Merge(t *testing.T) {
	// https://link.excalidraw.com/l/5SN39rbzyFm/4eTt0xgQYte
	a := Map{Edge{1, 2, 3}}
	b := Map{Edge{4, 2, 3}}

	actual := a.Merge(b)
	expected := Map{
		Edge{1, 2, 1},
		Edge{4, 3, 2},
	}
	diff := cmp.Diff(expected, actual)
	if diff != "" {
		t.Error(diff)
	}

	// Thanks!
	// https://github.com/steven-terrana/advent-of-code/blob/main/2023/day05/latex.md
	lines := util.GetLinesFromFile("input_small.txt")
	_, maps := parse(lines)
	merged := MergeAll(maps)
	expected = FromSensibleUnits([][3]int{
		{22, 0, 13},
		{29, 14, 14},
		{21, 15, 21},
		{68, 22, 25},
		{-25, 26, 43},
		{17, 44, 49},
		{-30, 50, 51},
		{-8, 52, 53},
		{31, 54, 58},
		{35, 59, 61},
		{-6, 62, 65},
		{31, 66, 68},
		{4, 69, 69},
		{-70, 70, 70},
		{3, 71, 81},
		{-36, 82, 91},
		{-32, 92, 92},
		{-25, 93, 97},
		{-31, 98, 98},
		{-80, 99, 99},
	})
	diff = cmp.Diff(expected, merged)
	if diff != "" {
		t.Error(diff)
	}
}
