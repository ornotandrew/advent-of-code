package day_06

import (
	"aoc2023/util"
	"math"
	"regexp"
	"strconv"
	"strings"
)

func Part1(lines []string) int {
	races := parse_1(lines)
	answer := 1
	for _, race := range races {
		minChargeTime, maxChargeTime := roots(race)
		numWays := maxChargeTime - minChargeTime + 1
		answer = answer * numWays
	}
	return answer
}

func Part2(lines []string) int {
	race := parse_2(lines)
	minChargeTime, maxChargeTime := roots(race)
	return maxChargeTime - minChargeTime + 1
}

type Race struct {
	Time     int
	Distance int
}

func parse_1(lines []string) []Race {
	times := util.GetNumbersBySeparator(lines[0], " ")
	distances := util.GetNumbersBySeparator(lines[1], " ")
	races := []Race{}
	for i, time := range times {
		races = append(races, Race{time, distances[i]})
	}
	return races
}

var numberRegex, _ = regexp.Compile("[0-9]([0-9 ]?)+")

func parse_2(lines []string) Race {
	rawTime := strings.ReplaceAll(numberRegex.FindString(lines[0]), " ", "")
	rawDistance := strings.ReplaceAll(numberRegex.FindString(lines[1]), " ", "")
	time, _ := strconv.Atoi(rawTime)
	distance, _ := strconv.Atoi(rawDistance)
	return Race{time, distance}
}

func roots(race Race) (int, int) {
	// chargeTime = x, distance = y
	// y = x * (race.Time - x)
	// y = -1 * x^2 + race.Time * x
	// if y > race.Distance then we can just solve for the roots
	// i.e. y = -1 * x^2 + race.Time * x - race.Distance
	a := float64(-1)
	b := float64(race.Time)
	c := float64(-race.Distance)
	r1 := (-b - math.Sqrt(math.Pow(b, 2)-4*a*c)) / (2 * a)
	r2 := (-b + math.Sqrt(math.Pow(b, 2)-4*a*c)) / (2 * a)
	minChargeTime, maxChargeTime := min(r1, r2), max(r1, r2)
	// if we match the results exactly, shift our value by 1 (we need to beat them)
	if minChargeTime == math.Ceil(minChargeTime) {
		minChargeTime += 1
	}
	if maxChargeTime == math.Floor(maxChargeTime) {
		maxChargeTime -= 1
	}
	return int(math.Ceil(minChargeTime)), int(math.Floor(maxChargeTime))
}
