package day_10

import (
	"fmt"
	"strings"
)

func Part1(b Board) int {
	start := b.StartingPosition()
	chain := b.FindChain(start)
	return (len(chain) / 2)
}

func Part2(b Board) int {
	return ScanLine(b, 'L')
}

func ScanLine(b Board, startReplacement rune) int {
	start := b.StartingPosition()
	chain := b.FindChain(start)

	b.ReplaceStartingPostition(start, startReplacement)
	chainMap := map[Pos]rune{}
	for _, pos := range chain {
		chainMap[pos] = b.Get(pos)
	}

	count := 0
	for row := 0; row < len(b); row++ {
		inside := false   // keep track of inside vs outsid
		lastCorner := 'x' // keep track of U shapes, which will not switch inside/outside
		for col := 0; col < len(b[0]); col++ {
			pos := Pos{row, col}
			if pipe, ok := chainMap[pos]; ok {
				switch pipe {
				case '|':
					inside = !inside
				case 'L':
					lastCorner = 'L'
				case 'F':
					lastCorner = 'F'
				case '7':
					if lastCorner != 'F' {
						inside = !inside
					}
					lastCorner = '7'
				case 'J':
					if lastCorner != 'L' {
						inside = !inside
					}
					lastCorner = 'J'
				}
			}

			_, isChain := chainMap[pos]
			if inside && !isChain {
				count++
			}
		}
	}
	return count
}

type Pos struct {
	Row int
	Col int
}

type Board []string

func (b Board) StartingPosition() Pos {
	startIdx := strings.Index(strings.Join(b, ""), "S")
	return Pos{startIdx / len(b[0]), startIdx % len(b[0])}
}

// just do it by eye
func (b Board) ReplaceStartingPostition(start Pos, value rune) {
	row := []rune(b[start.Row])
	row[start.Col] = value
	b[start.Row] = string(row)
}

func (b Board) GetValidSurroundingPipes(pos Pos) map[Pos][]rune {
	validConnections := map[Pos][]rune{}
	if pos.Row > 0 { // south connects
		validConnections[Pos{pos.Row - 1, pos.Col}] = []rune{'|', '7', 'F'}
	}
	if pos.Row < len(b)-1 { // north connects
		validConnections[Pos{pos.Row + 1, pos.Col}] = []rune{'|', 'L', 'J'}
	}
	if pos.Col > 0 { // east connects
		validConnections[Pos{pos.Row, pos.Col - 1}] = []rune{'-', 'L', 'F'}
	}
	if pos.Col < len(b[0])-1 { // west connects
		validConnections[Pos{pos.Row, pos.Col + 1}] = []rune{'-', 'J', '7'}
	}
	return validConnections
}

func (b Board) Get(pos Pos) rune {
	return rune(b[pos.Row][pos.Col])
}

func (b Board) GetConnectingPoints(pos Pos) []Pos {
	connecting := []Pos{}
	for validPos, validPipes := range b.GetValidSurroundingPipes(pos) {
		actualPipe := b.Get(validPos)
		for _, validPipe := range validPipes {
			if validPipe == actualPipe {
				connecting = append(connecting, validPos)
			}
		}
	}
	return connecting
}

func (b Board) FindChain(start Pos) []Pos {
	startingDirections := b.GetConnectingPoints(start)
	// just go in one direction
	chain := []Pos{startingDirections[0], start}
	prev := start
	current := startingDirections[1]
	length := 0
	for ; current != chain[0]; length++ {
		chain = append(chain, current)
		switch b.Get(current) {
		// | is a vertical pipe connecting north and south.
		// - is a horizontal pipe connecting east and west.
		// L is a 90-degree bend connecting north and east.
		// J is a 90-degree bend connecting north and west.
		// 7 is a 90-degree bend connecting south and west.
		// F is a 90-degree bend connecting south and east.
		// . is ground; there is no pipe in this tile.
		// S is the starting position of the animal; there is a pipe on this
		case '|':
			if prev.Row < current.Row {
				prev = current
				current = Pos{current.Row + 1, current.Col}
			} else if prev.Row > current.Row {
				prev = current
				current = Pos{current.Row - 1, current.Col}
			}
		case '-':
			if prev.Col < current.Col {
				prev = current
				current = Pos{current.Row, current.Col + 1}
			} else if prev.Col > current.Col {
				prev = current
				current = Pos{current.Row, current.Col - 1}
			}
		case 'L':
			if prev.Row < current.Row {
				prev = current
				current = Pos{current.Row, current.Col + 1}
			} else if prev.Row == current.Row {
				prev = current
				current = Pos{current.Row - 1, current.Col}
			}
		case 'J':
			if prev.Row < current.Row {
				prev = current
				current = Pos{current.Row, current.Col - 1}
			} else if prev.Row == current.Row {
				prev = current
				current = Pos{current.Row - 1, current.Col}
			}
		case '7':
			if prev.Row > current.Row {
				prev = current
				current = Pos{current.Row, current.Col - 1}
			} else if prev.Row == current.Row {
				prev = current
				current = Pos{current.Row + 1, current.Col}
			}
		case 'F':
			if prev.Row > current.Row {
				prev = current
				current = Pos{current.Row, current.Col + 1}
			} else if prev.Row == current.Row {
				prev = current
				current = Pos{current.Row + 1, current.Col}
			}
		default:
			panic(fmt.Sprintf("Found case: %v\n", b.Get(current)))
		}
	}
	return chain
}
