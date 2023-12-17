package main

import (
	"aoc2023/day_08"
	"aoc2023/util"
	"fmt"
	"os"
)

func main() {
	inputFname := os.Args[1]
	lines := util.GetLinesFromFile(inputFname)
	fmt.Println(day_08.Part2(lines))
}

