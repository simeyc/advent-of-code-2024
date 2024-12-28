package main

import (
	"fmt"
	"os"
	"slices"
	"strconv"
	"strings"
)

func part1(firstVals []int, secondVals []int) int {
	slices.Sort(firstVals)
	slices.Sort(secondVals)
	result := 0
	for i := range firstVals {
		if firstVals[i] > secondVals[i] {
			result += firstVals[i] - secondVals[i]
		} else {
			result += secondVals[i] - firstVals[i]
		}
	}
	return result
}

func part2(firstVals []int, secondVals []int) int {
	result := 0
	for _, val := range firstVals {
		for i := range secondVals {
			if secondVals[i] == val {
				result += val
			}
		}
	}
	return result
}

func main() {
	text, _ := os.ReadFile("./day1/input.txt")
	lines := strings.Split(string(text), "\n")
	firstVals := make([]int, len(lines))
	secondVals := make([]int, len(lines))
	for i, line := range lines {
		iSep := strings.Index(line, " ")
		firstVals[i], _ = strconv.Atoi(line[:iSep])
		secondVals[i], _ = strconv.Atoi(strings.Trim(line[iSep:], " "))
	}

	result := part1(firstVals, secondVals)
	fmt.Printf("Part 1 result: %d\n", result)


	result = part2(firstVals, secondVals)
	fmt.Printf("Part 2 result: %d\n", result)
}
