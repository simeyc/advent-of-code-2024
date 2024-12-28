package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func parseInput() [][]int {
	text, _ := os.ReadFile("./day2/input.txt")
	lines := strings.Split(string(text), "\n")
	levels := make([][]int, len(lines))
	for i := range lines {
		parts := strings.Split(lines[i], " ")
		levels[i] = make([]int, len(parts))
		for j := range parts {
			levels[i][j], _ = strconv.Atoi(parts[j])
		}
	}
	return levels
}

func isUnsafeStep(step int, asc bool) bool {
	return step == 0 || (asc && (step < 1 || step > 3)) || (!asc && (step > -1 || step < -3))
}

func part1(levels [][]int) int {
	result := 0
	for _, lvl := range levels {
		safe := true
		asc := lvl[1] > lvl[0]
		for i := 1; i < len(lvl); i++ {
			step := lvl[i] - lvl[i-1]
			if isUnsafeStep(step, asc) {
				safe = false
				break
			}
		}
		if safe {
			result += 1
		}
	}
	return result
}

func isSafeLevel(level []int) bool {
	asc := level[1] > level[0]
	for i := 1; i < len(level); i++ {
		step := level[i] - level[i-1]
		if isUnsafeStep(step, asc) {
			return false
		}
	}
	return true
}

func isSafeLevelWithDamping(level []int, index int) bool {
	lvl := level[:]
	dampedLvl := make([]int, len(level) - 1)
	for i := 0; i <= len(level); i++ {
		if !isSafeLevel(lvl) {
			if i == len(level) {
				return false
			}
			k := 0
			for j := range level {
				if j != i {
					dampedLvl[k] = level[j]
					k += 1
				}
			}
			lvl = dampedLvl
		}
	}
	return true
}

func part2(levels [][]int) int {
	result := 0
	for i := range levels {
		if isSafeLevelWithDamping(levels[i], i) {
			result++
		}
	}
	return result
}

func main() {
	levels := parseInput()

	result := part1(levels)
	fmt.Printf("Part 1 result: %d\n", result)

	result = part2(levels)
	fmt.Printf("Part 2 result: %d\n", result)
}