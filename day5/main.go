package main

import (
	"bufio"
	"fmt"
	"os"
	"slices"
	"strconv"
	"strings"
)

func part2(rules map[int][]int, updates [][]int) int {
	result := 0
	for _, update := range updates {
		for i := 1; i < len(update); i++ {
			iValueUnderTest := i
			for j := i-1; j >= 0; j-- {
				if slices.Contains(rules[update[iValueUnderTest]], update[j]) {
					// valueUnderTest must come first; swap.
					temp := update[j]
					update[j] = update[iValueUnderTest]
					update[iValueUnderTest] = temp
					iValueUnderTest = j
				}
			}
		}
		result += update[len(update) / 2]
	}
	return result
}

func main() {
	file, _ := os.Open("./day5/input.txt")
	defer file.Close()

	scanner := bufio.NewScanner(file)

	// Build map of rules; precedingPage: [list of succeeding pages].
	rules := make(map[int][]int)
	scanner.Scan()
	for text := scanner.Text(); text != ""; text = scanner.Text() {
		parts := strings.Split(text, "|")
		key, _ := strconv.Atoi(parts[0])
		val, _ := strconv.Atoi(parts[1])
		if _, ok := rules[key]; ok {
			rules[key] = append(rules[key], val)
		} else {
			rules[key] = []int{val}
		}
		scanner.Scan()
	}

	// Build list of updates.
	updates := [][]int{}
	for scanner.Scan() {
		parts := strings.Split(scanner.Text(), ",")
		update := make([]int, len(parts))
		for i := range parts {
			update[i], _ = strconv.Atoi(parts[i])
		}
		updates = append(updates, update)
	}

	// Iterate updates, determine which follow rules.
	result := 0
	incorrectUpdates := [][]int{}
updateLoop:
	for _, update := range updates {
		for i := 1; i < len(update); i++ {
			// Ensure no value in the rule for update[i] comes before it.
			for j := i; j >= 0; j-- {
				if slices.Contains(rules[update[i]], update[j]) {
					// Skip this update.
					incorrectUpdates = append(incorrectUpdates, update)
					continue updateLoop
				}
			}
		}
		// Add the middle value to the result.
		mid := len(update) / 2
		result += update[mid]
	}

	fmt.Printf("Part 1 result: %d\n", result)

	result = part2(rules, incorrectUpdates)
	fmt.Printf("Part 2 result: %d\n", result)
}