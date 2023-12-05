package main

import (
	"aoc2023/util"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Draw = [3]int
type Game = []Draw

func main() {
	inputFname := os.Args[1]
	lines := util.GetLinesFromFile(inputFname)
	fmt.Println(part1(lines, Draw{12, 13, 14}))
	fmt.Println(part2(lines))
}

func part1(lines []string, maxDraw Draw) int {
	games := parse(lines)
	score := 0
	for gameId, game := range games {
		for _, draw := range game {
			for i := 0; i < 3; i++ {
				if draw[i] > maxDraw[i] {
					goto Endgame
				}
			}
		}
		score += gameId + 1
	Endgame:
	}
	return score
}

func part2(lines []string) int64 {
	games := parse(lines)
	var score int64 = 0
	for _, game := range games {
		maxDraw := Draw{0, 0, 0}
		for _, draw := range game {
			for i := 0; i < 3; i++ {
				if draw[i] > maxDraw[i] {
					maxDraw[i] = draw[i]
				}
			}
		}
		score += int64(maxDraw[0]) * int64(maxDraw[1]) * int64(maxDraw[2])
	}
	return score
}

func parse(lines []string) []Game {
	var games [][][3]int
	for _, rawGame := range lines {
		draws := strings.Split(rawGame, ": ")
		draws = strings.Split(draws[1], "; ")
		var game Game
		for _, rawDraw := range draws {
			draw := Draw{0, 0, 0}
			rawColors := strings.Split(rawDraw, ", ")
			for _, rawColor := range rawColors {
				x := strings.Split(rawColor, " ")
				num, _ := strconv.Atoi(x[0])
				switch color := x[1]; color {
				case "red":
					draw[0] = num
				case "green":
					draw[1] = num
				case "blue":
					draw[2] = num
				}
			}
			game = append(game, draw)
		}
		games = append(games, game)
	}
	return games
}
