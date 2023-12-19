package day_11

import (
	"aoc2023/util"
	"math"
)

func Part1(lines []string) int {
	image := parse(lines)
	image = image.Expand(2)
	galaxies := []Pos{}
	for galaxy := range image.Galaxies {
		galaxies = append(galaxies, galaxy)
	}
	score := 0
	for _, pair := range util.UnorderedPairs(galaxies) {
		score += pair[0].Distance(pair[1])
	}
	return score
}

func Part2(lines []string) int {
	image := parse(lines)
	image = image.Expand(1000000)
	galaxies := []Pos{}
	for galaxy := range image.Galaxies {
		galaxies = append(galaxies, galaxy)
	}
	score := 0
	for _, pair := range util.UnorderedPairs(galaxies) {
		score += pair[0].Distance(pair[1])
	}
	return score
}

type Pos struct {
	Row int
	Col int
}

func (a Pos) Distance(b Pos) int {
	return int(math.Abs(float64(b.Row-a.Row)) + math.Abs(float64(b.Col-a.Col)))
}

type Image struct {
	Galaxies map[Pos]struct{}
	Rows     int
	Cols     int
}

func parse(lines []string) Image {
	galaxies := map[Pos]struct{}{}
	for rIdx, row := range lines {
		for cIdx, p := range row {
			if p == '#' {
				galaxies[Pos{rIdx, cIdx}] = struct{}{}
			}
		}
	}
	return Image{
		galaxies,
		len(lines),
		len(lines[0]),
	}
}

func (i Image) Expand(factor int) Image {
	emptyRowIdx := map[int]struct{}{}
	for row := 0; row < i.Rows; row++ {
		emptyRowIdx[row] = struct{}{}
	}

	emptyColIdx := map[int]struct{}{}
	for col := 0; col < i.Rows; col++ {
		emptyColIdx[col] = struct{}{}
	}

	for galaxy := range i.Galaxies {
		delete(emptyRowIdx, galaxy.Row)
		delete(emptyColIdx, galaxy.Col)
	}

	expandedGalaxies := map[Pos]struct{}{}
	rowShift := 0
	for row := 0; row < i.Rows; row++ {
		colShift := 0
		for col := 0; col < i.Cols; col++ {
			if _, ok := emptyColIdx[col]; ok {
				colShift += factor - 1
			}
			if _, ok := i.Galaxies[Pos{row, col}]; ok {
				expandedGalaxies[Pos{row + rowShift, col + colShift}] = struct{}{}
			}
		}
		if _, ok := emptyRowIdx[row]; ok {
			rowShift += factor - 1
		}
	}

	return Image{
		expandedGalaxies,
		0,
		0,
	}
}
