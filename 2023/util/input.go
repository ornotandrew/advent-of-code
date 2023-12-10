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
		num, _ := strconv.Atoi(strNum)
		numbers = append(numbers, num)
	}
	return numbers
}
