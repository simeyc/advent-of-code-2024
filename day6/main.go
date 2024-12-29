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

func main() {
	text, _ := os.ReadFile("./day6/input.txt")
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
	
	fmt.Printf("Part 1 result: %d\n", result)
}