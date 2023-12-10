package day_01

import (
	"fmt"
	"strconv"
	"strings"
)

func Part1(lines []string) int {
	return solve(lines, false)
}

func Part2(lines []string) int {
	return solve(lines, true)
}

func solve(lines []string, replace bool) int {
	replacements := []string{
		"one",
		"two",
		"three",
		"four",
		"five",
		"six",
		"seven",
		"eight",
		"nine",
	}
	count := 0
	for lineNumber := 0; lineNumber < len(lines); lineNumber++ {
		line := lines[lineNumber]
		var digitsInLine []int

		for position := 0; position < len(line); position++ {
			if digit, err := strconv.Atoi(line[position : position+1]); err == nil {
				digitsInLine = append(digitsInLine, digit)
				continue
			}
			if replace {
				for i, replacement := range replacements {
					if strings.HasPrefix(line[position:], replacement) {
						digitsInLine = append(digitsInLine, i+1)
						break
					}
				}
			}
		}
		leftAndRightSum, _ := strconv.Atoi(fmt.Sprintf("%d%d", digitsInLine[0], digitsInLine[len(digitsInLine)-1]))
		count += leftAndRightSum
	}
	return count
}
