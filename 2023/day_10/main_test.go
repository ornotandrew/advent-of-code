package day_10

import (
	"aoc2023/util"
	"testing"
	// "github.com/google/go-cmp/cmp"
)

func TestSolve_Part1(t *testing.T) {
	lines := util.GetLinesFromFile("input_small.txt")
	result := Part1(lines)
	expected := 4
	if result != expected {
		t.Errorf("Expected %d; Got %d", expected, result)
	}
}

func Test_ScanLine(t *testing.T) {
	type Test struct {
		fileName    string
		replacement rune
		expected    int
	}
	for _, test := range []Test{
		{"input_small.txt", 'F', 1},
		{"input_small_2.txt", 'F', 4},
		{"input_small_3.txt", 'F', 8},
		{"input_small_4.txt", '7', 10},
		{"input_small_5.txt", 'F', 4},
	} {
		lines := util.GetLinesFromFile(test.fileName)
		result := ScanLine(lines, test.replacement)
		if result != test.expected {
			t.Errorf("Case %s: Expected %d; Got %d\n", test.fileName, test.expected, result)
		}
	}
}
