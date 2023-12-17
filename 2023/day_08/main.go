package day_08

import (
	"aoc2023/util"
	"fmt"
	"regexp"
)

func Part1(lines []string) int {
	directions, nodeMap := parse(lines)
	currentNode, _ := nodeMap["AAA"]
	count := 0
	for currentNode.Id != "ZZZ" {
		currentNode = currentNode.Next(directions, count)
		count += 1
	}
	return count
}

func Part2(lines []string) int {
	// The way this question was phrased is EXTREMELY frustrating. I built most of a general solution before looking up the trick.
	// It turns out that
	//   1. Each starting node has a unique end node
	//   2. The lengths of the cycles containing the end nodes are equal to the distance from each start to each end node
	// This is not mentioned anywhere on the puzzle description, and you just need to realize it for yourself.
	// Basically, this means we can just work out the part 1 solution for each starting node and take the LCM.
	directions, nodeMap := parse(lines)
	distances := []int{}
	for _, node := range nodeMap {
		if node.Id[2] == 'A' {
			currentNode := node
			count := 0
			for currentNode.Id[2] != 'Z' {
				currentNode = currentNode.Next(directions, count)
				count += 1
			}
			distances = append(distances, count)
		}
	}

	return util.LCM(distances[0], distances[1], distances[1:]...)
}

type Node struct {
	Id    string
	Left  *Node
	Right *Node
}

func (n *Node) Next(directions string, position int) *Node {
	switch direction := directions[position%len(directions)]; direction {
	case 'L':
		return n.Left
	case 'R':
		return n.Right
	default:
		panic(1)
	}
}

func (n Node) String() string {
	return fmt.Sprintf("%s: (%s, %s)", n.Id, n.Left.Id, n.Right.Id)
}

var idRegex = regexp.MustCompile("[0-9A-Z]{2}[A-Z]")

func parse(lines []string) (string, map[string]*Node) {
	nodesById := map[string]*Node{}
	raw := [][3]string{}
	for _, line := range lines[2:] {
		matches := idRegex.FindAllStringSubmatch(line, -1)
		nodesById[matches[0][0]] = &Node{matches[0][0], nil, nil}
		raw = append(raw, [3]string{matches[0][0], matches[1][0], matches[2][0]})
	}

	for _, i := range raw {
		node, _ := nodesById[i[0]]
		left, _ := nodesById[i[1]]
		right, _ := nodesById[i[2]]
		node.Left = left
		node.Right = right
	}

	return lines[0], nodesById
}
