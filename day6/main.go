package main

import (
	"bytes"
	"fmt"
	"os"
)

// Walk as many steps as possible, from position pos on map areaMap in direction dir,
// until bumping the edge of the map or a '#' character. Updates pos and dir to new
// values. Paints 'X' chars into areaMap for each step. Returns number of steps taken
// and whether the edge of the map was reached.
func walk(areaMap *[][]byte, pos *[2]int, dir *rune) (int, bool) {
	steps := 0
	i := pos[0]
	j := pos[1]
	startDir := *dir
	maxI := len(*areaMap) - 1
	maxJ := len((*areaMap)[0]) - 1
	switch *dir {
	case '^':
		// Decrement line index until hitting edge or #.
		for ; i > 0 && (*areaMap)[i-1][pos[1]] != '#'; i-- {
			steps += 1
			(*areaMap)[i][pos[1]] = 'X'
		}
		if i > 0 {
			*dir = '>'
		}
	case 'v':
		// Increment line index until hitting edge or #.
		for ; i < maxI && (*areaMap)[i+1][pos[1]] != '#'; i++ {
			steps += 1
			(*areaMap)[i][pos[1]] = 'X'
		}
		if i < maxI {
			*dir = '<'
		}
	case '<':
		// Decrement line position until hitting edge or #.
		for ; j > 0 && (*areaMap)[pos[0]][j-1] != '#'; j-- {
			steps += 1
			(*areaMap)[pos[0]][j] = 'X'
		}
		if j > 0 {
			*dir = '^'
		}
	case '>':
		// Increment line position until hitting edge or #.
		for ; j < maxJ && (*areaMap)[pos[0]][j+1] != '#'; j++ {
			steps += 1
			(*areaMap)[pos[0]][j] = 'X'
		}
		if j < maxJ {
			*dir = 'v'
		}
	default:
		panic("Bad direction.")
	}
	*pos = [2]int{i,j}
	offMap := *dir == startDir
	if offMap {
		steps += 1
		(*areaMap)[i][j] = 'X'
	}
	return steps, offMap
}


// Walk as many steps as possible, from position pos on map areaMap in direction dir,
// until bumping the edge of the map or a '#' character. Updates pos and dir to new
// values. Paints '|'/'-'/'+' chars into areaMap for each step, marking hte route walked.
// Returns whether the edge of the map was reached.
func walk2(areaMap *[][]byte, pos *[2]int, dir *rune) bool {
	i := pos[0]
	j := pos[1]
	startDir := *dir
	maxI := len(*areaMap) - 1
	maxJ := len((*areaMap)[0]) - 1
	switch *dir {
	case '^':
		// Decrement line index until hitting edge or #.
		for ; i > 0 && (*areaMap)[i-1][pos[1]] != '#'; i-- {
			switch (*areaMap)[i][pos[1]] {
			case '.':
				(*areaMap)[i][pos[1]] = '|'
			case '-':
				(*areaMap)[i][pos[1]] = '+'
			default:
			}
		}
		if i > 0 {
			*dir = '>'
			(*areaMap)[i][pos[1]] = '+'
		}
	case 'v':
		// Increment line index until hitting edge or #.
		for ; i < maxI && (*areaMap)[i+1][pos[1]] != '#'; i++ {
			switch (*areaMap)[i][pos[1]] {
			case '.':
				(*areaMap)[i][pos[1]] = '|'
			case '-':
				(*areaMap)[i][pos[1]] = '+'
			default:
			}
		}
		if i < maxI {
			*dir = '<'
			(*areaMap)[i][pos[1]] = '+'
		}
	case '<':
		// Decrement line position until hitting edge or #.
		for ; j > 0 && (*areaMap)[pos[0]][j-1] != '#'; j-- {
			switch (*areaMap)[pos[0]][j] {
			case '.':
				(*areaMap)[pos[0]][j] = '-'
			case '|':
				(*areaMap)[pos[0]][j] = '+'
			default:
			}
		}
		if j > 0 {
			*dir = '^'
			(*areaMap)[pos[0]][j] = '+'
		}
	case '>':
		// Increment line position until hitting edge or #.
		for ; j < maxJ && (*areaMap)[pos[0]][j+1] != '#'; j++ {
			switch (*areaMap)[pos[0]][j] {
			case '.':
				(*areaMap)[pos[0]][j] = '-'
			case '|':
				(*areaMap)[pos[0]][j] = '+'
			default:
			}
		}
		if j < maxJ {
			*dir = 'v'
			(*areaMap)[pos[0]][j] = '+'
		}
	default:
		panic("Bad direction.")
	}
	*pos = [2]int{i,j}
	offMap := *dir == startDir
	if offMap {
		if (*dir == '^' || *dir == 'v') {
			if (*areaMap)[i][j] == '.' {
				(*areaMap)[i][j] = '|'
			} else if (*areaMap)[i][j] == '-' {
				(*areaMap)[i][j] = '+'
			}
		} else {
			if (*areaMap)[i][j] == '.' {
				(*areaMap)[i][j] = '-'
			} else if (*areaMap)[i][j] == '-' {
				(*areaMap)[i][j] = '+'
			}
		}
	}
	return offMap
}

func parseInput(filepath string) ([][]byte, [2]int, rune) {
	text, _ := os.ReadFile(filepath)
	areaMap := bytes.Split(text, []byte{'\n'})
	
	// Find starting position and direction.
	var pos [2]int
	var dir rune
posLoop:
	for i := range areaMap {
		for _, d := range []rune{'^','v','<','>'} {
			if j := bytes.IndexRune(areaMap[i], d); j >= 0 {
				pos = [2]int{i, j}
				dir = d
				break posLoop
			}
		} 
	}
	//fmt.Printf("Starting pos: (%d,%d); Direction: %v\n", pos[0], pos[1], string(dir))
	return areaMap, pos, dir
}

func part1() int {
	areaMap, pos, dir := parseInput("./day6/input.txt")

	offMap := false
	for !offMap {
		_, offMap = walk(&areaMap, &pos, &dir)
		//fmt.Printf("After %d steps: Pos=(%d,%d); Dir='%s'\n", steps, pos[0], pos[1], string(dir))
	}

	/*for i := range areaMap {
		fmt.Printf("%s\n", string(areaMap[i]))
	}*/

	result := 0
	for i := range areaMap {
		result += bytes.Count(areaMap[i], []byte{'X'})
	}

	return result
}

func main() {
	result := part1()	
	fmt.Printf("Part 1 result: %d\n", result)

	// Part 2
	// Walk the course, painting with '|'/'-'/'+' chars (don't overwrite starting pos).
	// If moving < and anywhere above me is a '+' under a '#', obstruction in the next slot creates a loop.
	// Same if moving ^ and anywhere to my right is a '+#',
	// or if moving > and anywhere below me is a '+' above a '#',
	// or if moving v and to the left is a '#+'.
	areaMap, pos, dir := parseInput("./day6/input_example.txt")
	offMap := false
	for !offMap {
		offMap = walk2(&areaMap, &pos, &dir)
	}

	for i := range areaMap {
		fmt.Printf("%s\n", string(areaMap[i]))
	}
}