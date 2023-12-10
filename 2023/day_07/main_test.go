package day_07

import (
	"aoc2023/util"
	"testing"

	"github.com/google/go-cmp/cmp"
)

func Test_TransformJokers(t *testing.T) {
	type Test struct {
		Given    Cards
		Expected Cards
	}
	for i, test := range []Test{
		{Cards{2, 2, 2, 2, 0}, Cards{2, 2, 2, 2, 2}},
		{Cards{2, 0, 0, 0, 0}, Cards{2, 2, 2, 2, 2}},
		{Cards{2, 2, 14, 14, 0}, Cards{2, 2, 14, 14, 14}},
	} {
		actual := test.Given.TransformJokers()
		diff := cmp.Diff(test.Expected, actual)
		if diff != "" {
			t.Errorf("Case %v: %v", i, diff)
		}
	}

}

func Test_parseWithoutJokers(t *testing.T) {
	lines := util.GetLinesFromFile("input_small.txt")
	actual := parse(lines, false)
	expected := Hands{
		{Cards{3, 2, 10, 3, 13}, 1, 765},
		{Cards{10, 5, 5, 11, 5}, 3, 684},
		{Cards{13, 13, 6, 7, 7}, 2, 28},
		{Cards{13, 10, 11, 11, 10}, 2, 220},
		{Cards{12, 12, 12, 11, 14}, 3, 483},
	}
	diff := cmp.Diff(expected, actual)
	if diff != "" {
		t.Errorf(diff)
	}
}

func Test_parseWithJokers(t *testing.T) {
	lines := util.GetLinesFromFile("input_small.txt")
	actual := parse(lines, true)
	expected := Hands{
		{Cards{3, 2, 10, 3, 13}, 1, 765},
		{Cards{10, 5, 5, 0, 5}, 5, 684},
		{Cards{13, 13, 6, 7, 7}, 2, 28},
		{Cards{13, 10, 0, 0, 10}, 5, 220},
		{Cards{12, 12, 12, 0, 14}, 5, 483},
	}
	diff := cmp.Diff(expected, actual)
	if diff != "" {
		t.Errorf(diff)
	}
}

func Test_Rank(t *testing.T) {
	actual := Hands{
		{Cards{3, 2, 10, 3, 13}, 1, 765},
		{Cards{10, 5, 5, 11, 5}, 3, 684},
		{Cards{13, 13, 6, 7, 7}, 2, 28},
		{Cards{13, 10, 11, 11, 10}, 2, 220},
		{Cards{12, 12, 12, 11, 14}, 3, 483},
	}
	actual.Rank()
	expected := Hands{
		{Cards{3, 2, 10, 3, 13}, 1, 765},
		{Cards{13, 10, 11, 11, 10}, 2, 220},
		{Cards{13, 13, 6, 7, 7}, 2, 28},
		{Cards{10, 5, 5, 11, 5}, 3, 684},
		{Cards{12, 12, 12, 11, 14}, 3, 483},
	}
	diff := cmp.Diff(expected, actual)
	if diff != "" {
		t.Errorf(diff)
	}
}

func TestSolve_Part1(t *testing.T) {
	lines := util.GetLinesFromFile("input_small.txt")
	result := Part1(lines)
	expected := 6440
	if result != expected {
		t.Errorf("Expected %d; Got %d", expected, result)
	}
}

func TestSolve_Part2(t *testing.T) {
	lines := util.GetLinesFromFile("input_small.txt")
	result := Part2(lines)
	expected := 5905
	if result != expected {
		t.Errorf("Expected %d; Got %d", expected, result)
	}
}
