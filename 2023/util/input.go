package util

import (
	"fmt"
	"os"
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
