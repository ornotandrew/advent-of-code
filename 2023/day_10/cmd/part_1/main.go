package main

import (
	"aoc2023/day_10"
	"aoc2023/util"
	"fmt"
	"os"
)

func main() {
	inputFname := os.Args[1]
	lines := util.GetLinesFromFile(inputFname)
	fmt.Println(day_10.Part1(lines))
}

