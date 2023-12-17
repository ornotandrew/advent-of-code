package day_09

import (
	"strconv"
	"strings"
)

func Part1(lines []string) int {
	sequences := parse(lines)
	score := 0
	for _, sequence := range sequences {
		layers := []Sequence{sequence}
		for layer := layers[0]; !layer.AllZero(); layer = layers[len(layers)-1] {
			nextLayer := Sequence{}
			for i := 1; i < len(layer); i++ {
				nextLayer = append(nextLayer, layer[i]-layer[i-1])
			}
			layers = append(layers, nextLayer)
		}

		layers[len(layers)-1] = append(layers[len(layers)-1], 0)
		for pos := len(layers) - 2; pos >= 0; pos-- {
			layers[pos] = append(layers[pos], layers[pos].Last()+layers[pos+1].Last())
		}

		score += layers[0].Last()
	}
	return score
}

func Part2(lines []string) int {
	sequences := parse(lines)
	score := 0
	for _, sequence := range sequences {
		layers := []Sequence{sequence}
		for layer := layers[0]; !layer.AllZero(); layer = layers[len(layers)-1] {
			nextLayer := Sequence{}
			for i := 1; i < len(layer); i++ {
				nextLayer = append(nextLayer, layer[i]-layer[i-1])
			}
			layers = append(layers, nextLayer)
		}

		layers[len(layers)-1] = append([]int{0}, layers[len(layers)-1]...)
		for pos := len(layers) - 2; pos >= 0; pos-- {
			layers[pos] = append([]int{layers[pos][0] - layers[pos+1][0]}, layers[pos]...)
		}

		score += layers[0][0]
	}
	return score
}

type Sequence []int

func (s Sequence) AllZero() bool {
	for _, val := range s {
		if val != 0 {
			return false
		}
	}
	return true
}

func (s Sequence) Last() int {
	return s[len(s)-1]
}

func parse(lines []string) []Sequence {
	sequences := []Sequence{}
	for _, line := range lines {
		sequence := Sequence{}
		for _, strNum := range strings.Split(line, " ") {
			num, _ := strconv.Atoi(strNum)
			sequence = append(sequence, num)
		}
		sequences = append(sequences, sequence)
	}
	return sequences
}
