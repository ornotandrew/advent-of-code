#!/bin/bash
day_number=$1
dir="day_${day_number}"

mkdir -p $dir/cmd/part_1/
mkdir -p $dir/cmd/part_2/
touch $dir/input.txt
touch $dir/input_small.txt

echo 'package main

import (
	"aoc2023/day_'${day_number}'"
	"aoc2023/util"
	"fmt"
	"os"
)

func main() {
	inputFname := os.Args[1]
	lines := util.GetLinesFromFile(inputFname)
	fmt.Println(day_'${day_number}'.Part1(lines))
}
' > $dir/cmd/part_1/main.go

echo 'package main

import (
	"aoc2023/day_'${day_number}'"
	"aoc2023/util"
	"fmt"
	"os"
)

func main() {
	inputFname := os.Args[1]
	lines := util.GetLinesFromFile(inputFname)
	fmt.Println(day_'${day_number}'.Part2(lines))
}
' > $dir/cmd/part_2/main.go

echo 'package day_'${day_number}'

func Part1(lines []string) int {
	return -1
}

func Part2(lines []string) int {
	return -1
}
' > $dir/main.go

echo 'package day_'${day_number}'

import (
	"aoc2023/util"
	"testing"
)

func TestSolve_Part1(t *testing.T) {
	lines := util.GetLinesFromFile("input_small.txt")
	result := Part1(lines)
	expected := -1
	if result != expected {
		t.Errorf("Expected %d; Got %d", expected, result)
	}
}

func TestSolve_Part2(t *testing.T) {
	lines := util.GetLinesFromFile("input_small.txt")
	result := Part2(lines)
	expected := -1
	if result != expected {
		t.Errorf("Expected %d; Got %d", expected, result)
	}
}
' > $dir/main_test.go


echo "Created $dir"
