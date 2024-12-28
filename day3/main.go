package main

import (
	"fmt"
	"os"
	"regexp"
	"strconv"
)

const pattern = `mul\(([\d]{1,3}),([\d]{1,3})\)`

func getMulTotal(re *regexp.Regexp, text []byte) int {
	submatches := re.FindAllSubmatch(text, -1)
	result := 0
	for _, matches := range submatches {
		x, _ := strconv.Atoi(string(matches[1]))
		y, _ := strconv.Atoi(string(matches[2]))
		result += x * y
	}
	return result
}

func part2(re *regexp.Regexp, text []byte) int {
	reDo := regexp.MustCompile(`do\(\)`)
	reDont := regexp.MustCompile(`don't\(\)`)
	// Find indices of "do()"s and "don't()"s.
	// Take slices of text in "do()" sections.
	// Call getMulTotal for each slice, sum results.
	doMatches := reDo.FindAllIndex(text, -1)
	dontMatches := reDont.FindAllIndex(text, -1)
	
	iDoMatches := -1
	iDontMatches := 0
	iDo := 0
	iDont := dontMatches[iDontMatches][0]
	result := 0
	for {
		// Find the next iDont which is after the current iDo.
		// If there are none, set iDont to end of text.
		for iDont < iDo {
			iDontMatches += 1
			if iDontMatches < len(dontMatches) {
				iDont = dontMatches[iDontMatches][0]
			} else {
				iDont = len(text)
			}
		}
		// Add result from slice in current "do" section.
		result += getMulTotal(re, text[iDo:iDont])
		// Find the next iDo after the current iDont.
		// If no more "do" sections, return result.
		for iDo <= iDont {
			iDoMatches += 1
			if iDoMatches < len(doMatches) {
				iDo = doMatches[iDoMatches][0]
			} else {
				return result
			}
		}
	}
}

func main() {
	text, _ := os.ReadFile("./day3/input.txt")
	//text = []byte("xmul(2,4)&mul[3,7]!^don't()_mul(5,5)+mul(32,64](mul(11,8)undo()?mul(8,5))")
	
	re := regexp.MustCompile(pattern)

	result := getMulTotal(re, text)
	fmt.Printf("Part 1 result: %d\n", result)

	result = part2(re, text)
	fmt.Printf("Part 2 result: %d\n", result)
}