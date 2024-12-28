package main

import (
	"fmt"
	"os"
	"strings"
)

func isXmasHoriz(line string, i int) int {
	result := 0
	if i >= 3 && line[i-3:i+1] == "SAMX" {
		result += 1
	}
	if len(line) >= i + 4 && line[i:i+4] == "XMAS" {
		result += 1
	}
	return result
}

func isXmasVert(lines []string, i int, j int) int {
	result := 0
	if i >= 3 && lines[i-1][j] == 'M' && lines[i-2][j] == 'A' && lines[i-3][j] == 'S' {
		result += 1
	}
	if len(lines) >= i + 4 && lines[i+1][j] == 'M' && lines[i+2][j] == 'A' && lines[i+3][j] == 'S' {
		result += 1
	}
	return result
}

func isXmasDiagUpward(lines []string, i int, j int) int {
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

func isXmasDiagDownward(lines []string, i int, j int) int {
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

func main() {
	text, _ := os.ReadFile("./day4/input.txt")
	//text = []byte("MMMSXXMASM\nMSAMXMSMSA\nAMXSXMAAMM\nMSAMASMSMX\nXMASAMXAMM\nXXAMMXXAMA\nSMSMSASXSS\nSAXAMASAAA\nMAMMMXMMMM\nMXMXAXMASX")
	lines := strings.Split(string(text), "\n")

	result := 0
	for i := range lines {
		for j := range lines[i] {
			if lines[i][j] == 'X' {
				result += isXmasHoriz(lines[i], j)
				result += isXmasVert(lines, i, j)
				result += isXmasDiagUpward(lines, i, j)
				result += isXmasDiagDownward(lines, i, j)
			}
		}
	}
	fmt.Printf("Part 1 result: %d\n", result)
}