package main

import (
	"aoc2023/util"
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"
)

func main() {
	inputFname := os.Args[1]
	lines := util.GetLinesFromFile(inputFname)
	fmt.Println(part1(lines))
	fmt.Println(part2(lines))
}

func part1(lines []string) int {
	numbers := findNumbers(lines)
	symbols := findSymbols(lines)
	score := 0

	for _, num := range numbers {
		if numberIsAdjacentToSymbol(num, symbols) {
			score += num.Num
		}
	}

	return score
}

func part2(lines []string) int {
	numbers := findNumbers(lines)
	symbols := findSymbols(lines)
	stars := []Symbol{}
	for _, symbol := range symbols {
		if symbol.Value == '*' {
			stars = append(stars, symbol)
		}
	}

	score := 0
	for _, star := range stars {
		adjacentNumbers := getAdjacentNumbers(star, numbers)
		if len(adjacentNumbers) == 2 {
			score += (adjacentNumbers[0].Num * adjacentNumbers[1].Num)
		}
	}
	return score
}

type Number struct {
	Row      int
	StartCol int
	EndCol   int
	Num      int
}

func (n Number) Contains(position Position) bool {
	return position.Row == n.Row && (position.Col >= n.StartCol && position.Col <= n.EndCol)
}

var numberRegex, _ = regexp.Compile("[0-9]+")

func findNumbers(lines []string) []Number {
	rows := len(lines)
	cols := len(lines[0])

	joinedLines := strings.Join(lines, "")
	locations := numberRegex.FindAllStringIndex(joinedLines, -1)
	matches := []Number{}
	for _, loc := range locations {
		row := loc[0] / rows
		num, _ := strconv.Atoi(joinedLines[loc[0]:loc[1]])
		matches = append(matches, Number{
			row,
			loc[0] % cols,
			(loc[1] - 1) % cols,
			num,
		})
	}

	return matches
}

type Position struct {
	Row int
	Col int
}

type Symbol struct {
	Row   int
	Col   int
	Value rune
}

var symbolRegex, _ = regexp.Compile("[^0-9\\.]")

func findSymbols(lines []string) []Symbol {
	rows := len(lines)
	cols := len(lines[0])

	joinedLines := strings.Join(lines, "")
	locations := symbolRegex.FindAllStringIndex(joinedLines, -1)
	matches := []Symbol{}
	for _, loc := range locations {
		row := loc[0] / rows
		matches = append(matches, Symbol{
			row,
			loc[0] % cols,
			rune(joinedLines[loc[0]]),
		})
	}

	return matches
}

func get9(pos Position) []Position {
	result := []Position{}
	for _, dRow := range []int{-1, 0, 1} {
		for _, dCol := range []int{-1, 0, 1} {
			result = append(result, Position{pos.Row + dRow, pos.Col + dCol})
		}
	}
	return result
}

func numberIsAdjacentToSymbol(num Number, symbols []Symbol) bool {
	for col := num.StartCol; col <= num.EndCol; col++ {
		for _, pos := range get9(Position{num.Row, col}) {
			for _, symbol := range symbols {
				if symbol.Row == pos.Row && symbol.Col == pos.Col {
					return true
				}
			}
		}

	}
	return false

}

func getAdjacentNumbers(symbol Symbol, numbers []Number) []Number {
	adjacentNumbers := []Number{}
	surroundingPositions := get9(Position{symbol.Row, symbol.Col})
	for _, num := range numbers {
		for _, pos := range surroundingPositions {
			if num.Contains(pos) {
				adjacentNumbers = append(adjacentNumbers, num)
				break
			}
		}
	}
	return adjacentNumbers
}
