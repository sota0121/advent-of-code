package main

import (
	"bufio"
	"fmt"
	"os"
)

type Direction string

const (
	Vertical      Direction = "Vertical"
	Horizontal    Direction = "Horizontal"
	DiagonalRight Direction = "DiagonalRight"
	DiagonalLeft  Direction = "DiagonalLeft"
)

func LogDebug(format string, args ...interface{}) {
	if DEBUG {
		fmt.Printf(format, args...)
	}
}

var (
	DEBUG = true
	KW    = "XMAS"
)

func main() {
	// read data from stdin
	grid := readStdinAsMatrix()

	// scan the grid for the word
	count := scan(grid, KW, Vertical)
	count += scan(grid, KW, Horizontal)
	count += scan(grid, KW, DiagonalRight)
	count += scan(grid, KW, DiagonalLeft)

	// print the result
	fmt.Printf("total occurrences of %s: %d\n", KW, count)
}

func readStdinAsMatrix() [][]rune {
	scanner := bufio.NewScanner(os.Stdin)
	grid := make([][]rune, 0)
	for scanner.Scan() {
		line := scanner.Text()
		grid = append(grid, []rune(line))
	}
	return grid
}

func reverse(slice []rune) []rune {
	// reverse the slice
	// use two pointers
	// one from the beginning
	// one from the end
	// swap the values
	// move the pointers
	// until they meet in the middle
	// or until the first pointer is greater than the second pointer
	// return the reversed slice

	// get copy of the slice
	reversed := make([]rune, len(slice))
	copy(reversed, slice)
	for i, j := 0, len(slice)-1; i < j; i, j = i+1, j-1 {
		reversed[i], reversed[j] = reversed[j], reversed[i] // swap
	}
	return reversed
}

// scan scans the grid for the word
// returns the number of times the word is found
func scan(grid [][]rune, word string, dir Direction) int {
	// set variables
	windowSize := len(word)
	rows := len(grid)
	cols := 0
	if rows > 0 {
		cols = len(grid[0])
	}
	count := 0

	// check if the grid is empty
	if rows == 0 || cols == 0 {
		return count
	}

	// ------------------------------
	// core process
	// ------------------------------
	// scan vertically
	switch dir {
	case Vertical:
		for c := 0; c < cols; c++ {
			for r := 0; r < rows-windowSize+1; r++ {
				segment := make([]rune, windowSize)
				for i := 0; i < windowSize; i++ {
					segment[i] = grid[r+i][c]
				}
				// forward
				if string(segment) == word {
					count++
					LogDebug("%s: c=%d, Forward: %s, count: %d\n", dir, c, string(segment), count)
				}
				// backward
				revsegment := reverse(segment)
				if string(revsegment) == word {
					count++
					LogDebug("%s: c=%d, Backward: %s, count: %d\n", dir, c, string(revsegment), count)
				}
			}
		}
	case Horizontal:
		for r := 0; r < rows; r++ {
			for c := 0; c < cols-windowSize+1; c++ {
				segment := make([]rune, windowSize)
				for i := 0; i < windowSize; i++ {
					segment[i] = grid[r][c+i]
				}
				// forward
				if string(segment) == word {
					count++
					LogDebug("%s: c=%d, Forward: %s, count: %d\n", dir, c, string(segment), count)
				}
				// backward
				revsegment := reverse(segment)
				if string(revsegment) == word {
					count++
					LogDebug("%s: c=%d, Backward: %s, count: %d\n", dir, c, string(revsegment), count)
				}
			}
		}
	case DiagonalRight:
		fmt.Println("DiagonalRight: not implemented")
	case DiagonalLeft:
		fmt.Println("DiagonalLeft: not implemented")
	}
	return count
}
