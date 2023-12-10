package day_05

import (
	"aoc2023/util"
	"sort"
)

func Part1(lines []string) int {
	seeds, maps := parse(lines)
	minVal := -1
	for _, currentVal := range seeds {
		for _, map_ := range maps {
			currentVal = map_.Next(currentVal)
		}
		if minVal == -1 || currentVal < minVal {
			minVal = currentVal
		}
	}
	return minVal
}

func Part2(lines []string) int {
	seedRanges, maps := parse(lines)
	merged := MergeAll(maps)

	// search the solution space by starting with the lowest destination edge until
	// we find a matching seed
	seedBuckets := constructSeedBuckets(seedRanges)
	smallestSeed := findSmallestSeed(seedBuckets, merged)
	return merged.Next(smallestSeed)
}

func findSmallestSeed(seedBuckets SeedBuckets, merged Map) int {
	seedBuckets.SortAsc()
	merged.SortByDestAsc()
	for _, edge := range merged {
		// find the lowest seed that exists inside the src range
		for _, seedBucket := range seedBuckets {
			if seedBucket[0] > edge.MaxSrcValue() {
				// the seed bucket is too high, which means our edge is too low
				break
			}
			if seedBucket[0]+seedBucket[1] < edge[1] {
				// the seed bucket is too low, which means we can use the identity mapping
				continue
			}
			// otherwise, there is an overlap - find the smallest seed value within the edge
			return max(seedBucket[0], edge[1])
		}
	}
	return -1
}

type Seeds []int
type SeedBuckets [][2]int // {start, rangeCount}
func (sb SeedBuckets) SortAsc() {
	sort.Slice(sb, func(i, j int) bool {
		return sb[i][0] < sb[j][0]
	})
}

func parse(lines []string) (Seeds, []Map) {
	maps := []Map{}
	seeds := util.GetNumbersBySeparator(lines[0][7:], " ")

	currentMap := Map{}
	lineNum := 3
	for lineNum < len(lines) {
		if lines[lineNum] == "" {
			maps = append(maps, currentMap)
			currentMap = Map{}
			lineNum += 2
			continue
		}

		lineValues := util.GetNumbersBySeparator(lines[lineNum], " ")
		currentMap = append(currentMap, [3]int(lineValues))
		lineNum += 1
	}
	maps = append(maps, currentMap)

	return seeds, maps
}

func constructSeedBuckets(seeds Seeds) SeedBuckets {
	seedBuckets := [][2]int{}
	for i := 0; i < len(seeds); i += 2 {
		seedBuckets = append(seedBuckets, [2]int{seeds[i], seeds[i+1]})
	}
	return seedBuckets
}
