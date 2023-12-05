#!/bin/bash
day_number=$1
dir=day-$day_number

mkdir $dir
touch $dir/input.txt

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
}

func part2(lines []string) int {
}
' > $dir/day_${day_number}.go

echo 'package main

import "testing"

func TestSolve_Part1(t *testing.T) {
	answer := part1()
	expected := 1
	if answer != expected {
		t.Errorf("Expected %d; Got %d", expected, answer)
	}
}

func TestSolve_Part2(t *testing.T) {
	answer := part2()
	expected := 1
	if answer != expected {
		t.Errorf("Expected %d; Got %d", expected, answer)
	}
}
' > $dir/day_${day_number}_test.go


echo "Created $dir"
