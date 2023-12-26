package day_12

import (
	"strconv"
	"strings"
)

func Part1(lines []string) int {
	rows := parse(lines)
	score := 0
	for _, row := range rows {
		score += solve(&State{}, row, 0, 0, 0)
	}
	return score
}

func Part2(lines []string) int {
	rows := []Row{}
	for _, parsed := range parse(lines) {
		rows = append(rows, parsed.Expand())
	}
	score := 0
	for _, row := range rows {
		score += solve(&State{}, row, 0, 0, 0)
	}
	return score
}

type State map[[3]int]int
type Row struct {
	Springs string
	Groups  []int
}

func (r Row) Expand() Row {
	newRow := Row{"", []int{}}
	for i := 0; i < 5; i++ {
		newRow.Springs = newRow.Springs + r.Springs
		if i != 4 {
			newRow.Springs = newRow.Springs + "?"
		}
	}
	for i := 0; i < 5; i++ {
		newRow.Groups = append(newRow.Groups, r.Groups...)
	}
	return newRow
}

func parse(lines []string) []Row {
	rows := []Row{}
	for _, line := range lines {
		raw := strings.Split(line, " ")
		strNums := strings.Split(raw[1], ",")
		groups := []int{}
		for _, strNum := range strNums {
			num, _ := strconv.Atoi(strNum)
			groups = append(groups, num)
		}
		rows = append(rows, Row{raw[0], groups})
	}
	return rows
}

func solve(validRowsByState *State, row Row, pos, groupIdx, groupLen int) int {
	state := [3]int{pos, groupIdx, groupLen}
	if validRows, ok := (*validRowsByState)[state]; ok {
		return validRows
	}

	// check final states
	if pos == len(row.Springs) {
		if groupIdx == len(row.Groups) && groupLen == 0 {
			return 1
		}
		// we're still on the last job, because the last character was a #
		if groupIdx == len(row.Groups)-1 && groupLen == row.Groups[groupIdx] {
			return 1
		}
		return 0
	}

	// look into the future and figure out how many valid rows will result from the current state
	validRows := 0
	for _, option := range []rune{'.', '#'} {
		currentChar := rune(row.Springs[pos])
		if currentChar != '?' && currentChar != option {
			// this isn't even a valid thing to check
			continue
		}
		switch option {
		case '.':
			if groupLen == 0 { // we aren't in a group, so do nothing
				validRows += solve(validRowsByState, row, pos+1, groupIdx, 0)
			} else { // we're inside a group, which means we're ending it and will increment CurrentGroup
				if groupIdx >= len(row.Groups) {
					// we can't possibly increment here, since we'd have too many groups
					continue
				}
				if groupLen != row.Groups[groupIdx] {
					// group length doesn't match - not valid
					continue
				}
				validRows += solve(validRowsByState, row, pos+1, groupIdx+1, 0)
			}
		case '#':
			validRows += solve(validRowsByState, row, pos+1, groupIdx, groupLen+1)
		}
	}

	(*validRowsByState)[state] = validRows

	return validRows
}
