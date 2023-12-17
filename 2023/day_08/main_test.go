package day_08

import (
	"aoc2023/util"
	"testing"

	"github.com/google/go-cmp/cmp"
)

func Test_parse(t *testing.T) {
	lines := util.GetLinesFromFile("input_small.txt")
	directions, nodeMap := parse(lines)

	expectedDirections := "RL"
	if directions != expectedDirections {
		t.Errorf("Expected %v but fot %v", expectedDirections, directions)
	}

	aaa := Node{"AAA", nil, nil}
	bbb := Node{"BBB", nil, nil}
	ccc := Node{"CCC", nil, nil}
	ddd := Node{"DDD", nil, nil}
	eee := Node{"EEE", nil, nil}
	ggg := Node{"GGG", nil, nil}
	zzz := Node{"ZZZ", nil, nil}
	aaa.Left = &bbb
	aaa.Right = &ccc
	bbb.Left = &ddd
	bbb.Right = &eee
	ccc.Left = &zzz
	ccc.Right = &ggg
	ddd.Left = &ddd
	ddd.Right = &ddd
	eee.Left = &eee
	eee.Right = &eee
	ggg.Left = &ggg
	ggg.Right = &ggg
	zzz.Left = &zzz
	zzz.Right = &zzz

	diff := cmp.Diff(aaa, *nodeMap["AAA"])
	if diff != "" {
		t.Error(diff)
	}
}

func TestSolve_Part1(t *testing.T) {
	lines := util.GetLinesFromFile("input_small.txt")
	result := Part1(lines)
	expected := 2
	if result != expected {
		t.Errorf("Expected %d; Got %d", expected, result)
	}

	lines = util.GetLinesFromFile("input_small_2.txt")
	result = Part1(lines)
	expected = 6
	if result != expected {
		t.Errorf("Expected %d; Got %d", expected, result)
	}
}

func TestSolve_Part2(t *testing.T) {
	lines := util.GetLinesFromFile("input_small_3.txt")
	result := Part2(lines)
	expected := 6
	if result != expected {
		t.Errorf("Expected %d; Got %d", expected, result)
	}
}
