package day_07

import (
	"sort"
	"strconv"
	"strings"
)

func Part1(lines []string) int {
	hands := parse(lines, false)
	hands.Rank()
	winnings := 0
	for i, hand := range hands {
		winnings += hand.Score * (i + 1)
	}
	return winnings
}

func Part2(lines []string) int {
	hands := parse(lines, true)
	hands.Rank()
	winnings := 0
	for i, hand := range hands {
		winnings += hand.Score * (i + 1)
	}
	return winnings
}

func cardOrder(r rune, useJokers bool) int {
	switch r {
	case '2':
		return 2
	case '3':
		return 3
	case '4':
		return 4
	case '5':
		return 5
	case '6':
		return 6
	case '7':
		return 7
	case '8':
		return 8
	case '9':
		return 9
	case 'T':
		return 10
	case 'J':
		if useJokers {
			return 0
		}
		return 11
	case 'Q':
		return 12
	case 'K':
		return 13
	case 'A':
		return 14
	default:
		panic(1)
	}
}

type Cards [5]int

func (cards Cards) Category() int {
	counts := map[int]int{}
	for _, card := range cards {
		if _, ok := counts[card]; !ok {
			counts[card] = 0
		}
		counts[card]++
	}

	countValues := []int{}
	for _, value := range counts {
		countValues = append(countValues, value)
	}
	sort.Ints(countValues)
	handSignature := ""
	for _, i := range countValues {
		handSignature += strconv.Itoa(i)
	}

	switch handSignature {
	case "5":
		return 6 // five of a kind
	case "14":
		return 5 // four of a kind
	case "23":
		return 4 // full house
	case "113":
		return 3 // three of a kind
	case "122":
		return 2 // two pair
	case "1112":
		return 1 // one pair
	case "11111":
		return 0 // high card
	default:
		panic(1)
	}
}

func (cards Cards) TransformJokers() Cards {
	// The correct transformation will always transform all jokers to the same value
	transformedHands := Hands{}
	for _, transformation := range []int{2, 3, 4, 5, 6, 7, 8, 9, 10, 12, 13, 14} {
		transformedCards := cards.Replace(0, transformation)
		transformedHands = append(transformedHands, Hand{
			transformedCards,
			transformedCards.Category(),
			-1,
		})
	}
	transformedHands.Rank()
	return transformedHands[len(transformedHands)-1].Cards
}

func (cards Cards) Replace(src, dest int) Cards {
	newCards := Cards{}
	for i := 0; i < 5; i++ {
		if cards[i] == src {
			newCards[i] = dest
			continue
		}
		newCards[i] = cards[i]
	}
	return newCards
}

type Hand struct {
	Cards    Cards
	Category int
	Score    int
}

type Hands []Hand

func (hands Hands) Rank() {
	sort.Slice(hands, func(i, j int) bool {
		if hands[i].Category == hands[j].Category {
			for cardIdx := 0; cardIdx < 5; cardIdx++ {
				if hands[i].Cards[cardIdx] == hands[j].Cards[cardIdx] {
					continue
				}
				return hands[i].Cards[cardIdx] < hands[j].Cards[cardIdx]
			}
		}
		return hands[i].Category < hands[j].Category
	})
}

func parse(lines []string, useJokers bool) Hands {
	hands := Hands{}
	for _, line := range lines {
		rawValues := strings.Split(line, " ")
		score, _ := strconv.Atoi(rawValues[1])
		cards := Cards{}
		for i, card := range rawValues[0] {
			cards[i] = cardOrder(card, useJokers)
		}
		category := 0
		if useJokers {
			category = cards.TransformJokers().Category()
		} else {
			category = cards.Category()
		}
		hands = append(hands, Hand{cards, category, score})

	}
	return hands
}
