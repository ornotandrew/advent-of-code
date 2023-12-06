package main

import (
	"aoc2023/util"
	"fmt"
	"math"
	"os"
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
	cards := parse(lines)
	score := 0

	for _, card := range cards {
		count := card.Count()
		if count > 0 {
			score += int(math.Pow(2, float64(count)-1))
		}
	}
	return score
}

func part2(lines []string) int {
	cards := parse(lines)
	counts := []int{}
	copies := []int{}

	for _, card := range cards {
		// we start off with 1 copy of each card
		copies = append(copies, 1)
		counts = append(counts, card.Count())
	}

	for cardNumber := 0; cardNumber < len(counts); cardNumber++ {
		for cardOffset := 1; cardOffset < counts[cardNumber]+1; cardOffset++ {
			copies[cardNumber+cardOffset] += copies[cardNumber]
		}
	}

	copiesCount := 0
	for _, numCopies := range copies {
		copiesCount += numCopies
	}

	return copiesCount
}

type Card struct {
	WinningNumbers []int
	Numbers        map[int]struct{}
}

func (c Card) Count() int {
	count := 0
	for _, winningNumber := range c.WinningNumbers {
		if _, ok := c.Numbers[winningNumber]; ok {
			count++
		}
	}
	return count
}

func parse(lines []string) []Card {
	cards := []Card{}

	for _, line := range lines {
		raw := strings.Split(line, " | ")
		card := Card{
			[]int{},
			map[int]struct{}{},
		}
		for _, strNum := range strings.Split(raw[0][strings.Index(raw[0], ": ")+1:], " ") {
			if strNum == "" {
				continue
			}
			num, _ := strconv.Atoi(strNum)
			card.WinningNumbers = append(card.WinningNumbers, num)
		}
		for _, strNum := range strings.Split(raw[1], " ") {
			if strNum == "" {
				continue
			}
			num, _ := strconv.Atoi(strNum)
			card.Numbers[num] = struct{}{}
		}
		cards = append(cards, card)
	}

	return cards
}
