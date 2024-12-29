package main

import (
	"fmt"
	"os"
	"strings"
)

func countXmasHoriz(line string, i int) int {
	result := 0
	if i >= 3 && line[i-3:i+1] == "SAMX" {
		result += 1
	}
	if len(line) >= i + 4 && line[i:i+4] == "XMAS" {
		result += 1
	}
	return result
}

func countXmasVert(lines []string, i int, j int) int {
	result := 0
	if i >= 3 && lines[i-1][j] == 'M' && lines[i-2][j] == 'A' && lines[i-3][j] == 'S' {
		result += 1
	}
	if len(lines) >= i + 4 && lines[i+1][j] == 'M' && lines[i+2][j] == 'A' && lines[i+3][j] == 'S' {
		result += 1
	}
	return result
}

func countXmasDiagUpward(lines []string, i int, j int) int {
	result := 0
	if i < 3 {
		return 0
	}
	if j >= 3 && lines[i-1][j-1] == 'M' && lines[i-2][j-2] == 'A' && lines[i-3][j-3] == 'S' {
		result += 1 // diagonally upwards and backwards
	}
	if j+3 < len(lines) && lines[i-1][j+1] == 'M' && lines[i-2][j+2] == 'A' && lines[i-3][j+3] == 'S' {
		result += 1 // diagonally upwards and forwards
	}
	return result
}

func countXmasDiagDownward(lines []string, i int, j int) int {
	result := 0
	if i + 3 >= len(lines) {
		return 0
	}
	if j >= 3 && lines[i+1][j-1] == 'M' && lines[i+2][j-2] == 'A' && lines[i+3][j-3] == 'S' {
		result += 1 // diagonally downwards and backwards
	}
	if j+3 < len(lines) && lines[i+1][j+1] == 'M' && lines[i+2][j+2] == 'A' && lines[i+3][j+3] == 'S' {
		result += 1 // diagonally downwards and forwards
	}
	return result
}

func part1(lines []string) int {
	result := 0
	for i := range lines {
		for j := range lines[i] {
			if lines[i][j] == 'X' {
				result += countXmasHoriz(lines[i], j)
				result += countXmasVert(lines, i, j)
				result += countXmasDiagUpward(lines, i, j)
				result += countXmasDiagDownward(lines, i, j)
			}
		}
	}
	return result
}

type MSCoords struct {
	mCoords [2][2]int
	sCoords [2][2]int
}

// i and j are coords of an "A", returns whether that "A" is the center of an "X-MAS"
func isXmas(lines []string, i int, j int) bool {
	possibleCoords := []MSCoords{
		{									 // M-M
			[2][2]int{{i-1,j-1}, {i-1,j+1}}, // -A-
			[2][2]int{{i+1,j+1}, {i+1,j-1}}, // S-S
		},
		{									 // M-S
			[2][2]int{{i-1,j-1}, {i+1,j-1}}, // -A-
			[2][2]int{{i-1,j+1}, {i+1,j+1}}, // M-S
		},
		{									 // S-S
			[2][2]int{{i+1,j-1}, {i+1,j+1}}, // -A-
			[2][2]int{{i-1,j-1}, {i-1,j+1}}, // M-M
		}, 
		{									 // S-M
			[2][2]int{{i-1,j+1}, {i+1,j+1}}, // -A-
			[2][2]int{{i-1,j-1}, {i+1,j-1}}, // S-M
		}, 
	}

	for _, coords := range possibleCoords {
		if lines[coords.mCoords[0][0]][coords.mCoords[0][1]] == 'M' &&
				lines[coords.mCoords[1][0]][coords.mCoords[1][1]] == 'M' &&
				lines[coords.sCoords[0][0]][coords.sCoords[0][1]] == 'S' &&
				lines[coords.sCoords[1][0]][coords.sCoords[1][1]] == 'S' {
			return true
		}
	}
	return false
}

func part2(lines []string) int {
	result := 0
	for i := 1; i < len(lines) - 1; i++ {
		for j := 1; j < len(lines[0]) - 1; j++ {
			if lines[i][j] == 'A' && isXmas(lines, i, j) {
				result += 1
			}
		}
	}
	return result
}

func main() {
	text, _ := os.ReadFile("./day4/input.txt")
	//text = []byte("MMMSXXMASM\nMSAMXMSMSA\nAMXSXMAAMM\nMSAMASMSMX\nXMASAMXAMM\nXXAMMXXAMA\nSMSMSASXSS\nSAXAMASAAA\nMAMMMXMMMM\nMXMXAXMASX")
	lines := strings.Split(string(text), "\n")

	result := part1(lines)
	fmt.Printf("Part 1 result: %d\n", result)

	result = part2(lines)
	fmt.Printf("Part 2 result: %d\n", result)
}