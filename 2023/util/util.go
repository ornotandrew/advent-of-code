package util

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func GetLinesFromFile(fName string) []string {
	data, err := os.ReadFile(fName)
	if err != nil {
		msg, _ := fmt.Printf("Couldn't read file at %s", fName)
		panic(msg)
	}

	return strings.Split(strings.TrimSpace(string(data)), "\n")
}

func GetNumbersBySeparator(str, separator string) []int {
	numbers := []int{}
	for _, strNum := range strings.Split(str, " ") {
		num, err := strconv.Atoi(strNum)
		if err != nil {
			continue
		}
		numbers = append(numbers, num)
	}
	return numbers
}

// greatest common divisor (GCD) via Euclidean algorithm
func GCD(a, b int) int {
	for b != 0 {
		t := b
		b = a % b
		a = t
	}
	return a
}

// find Least Common Multiple (LCM) via GCD
func LCM(a, b int, integers ...int) int {
	result := a * b / GCD(a, b)

	for i := 0; i < len(integers); i++ {
		result = LCM(result, integers[i])
	}

	return result
}
