#!/bin/bash
day_number=$1
dir=day-$day_number

mkdir $dir
touch $dir/input.txt
touch $dir/input_small.txt

echo 'package main

import (
	"aoc2023/util"
	"fmt"
	"os"
)

func main() {
	inputFname := os.Args[1]
	lines := util.GetLinesFromFile(inputFname)
	fmt.Println(part1(lines))
	fmt.Println(part2(lines))
}

func part1(lines []string) int {
	return -1
}

func part2(lines []string) int {
	return -1
}
' > $dir/day_${day_number}.go

echo 'package main

import (
	"aoc2023/util"
	"testing"
)

func TestSolve_Part1(t *testing.T) {
	lines := util.GetLinesFromFile("input_small.txt")
	result := part1(lines)
	expected := -1
	if result != expected {
		t.Errorf("Expected %d; Got %d", expected, result)
	}
}

func TestSolve_Part2(t *testing.T) {
	lines := util.GetLinesFromFile("input_small.txt")
	result := part2(lines)
	expected := -1
	if result != expected {
		t.Errorf("Expected %d; Got %d", expected, result)
	}
}
' > $dir/day_${day_number}_test.go


echo "Created $dir"
