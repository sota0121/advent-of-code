package main

import (
	"bufio"
	"encoding/csv"
	"fmt"
	"os"
	"strconv"
)

var DEBUG = false

func main() {
	// note: unavailable to use the seek position back to top of the file for stdin
	// that's why we don't use the following code
	// ---
	// // scan and allocate memory for the input data
	// scanner := bufio.NewScanner(os.Stdin)
	// numOfLines := 0
	// for scanner.Scan() {
	// 	numOfLines++
	// }
	// if err := scanner.Err(); err != nil {
	// 	panic(err)
	// }
	// left, right := make([]int, 0, numOfLines), make([]int, 0, numOfLines)

	// // reset the seek position of the input data
	// if _, err := os.Stdin.Seek(0, 0); err != nil {
	// 	panic(err)
	// }
	// ---
	// ---
	// ---

	// read data from stdin (tsv)
	left, right := make([]int, 0), make([]int, 0)
	reader := csv.NewReader(bufio.NewReader(os.Stdin))
	// reader.Comma = '\t'
	records, err := reader.ReadAll()
	if err != nil {
		fmt.Println("Error:", err)
		panic(err)
	}

	lineNum := 0

	for _, record := range records {
		// check
		if len(record) != 2 {
			panic("invalid input data")
		}

		// parse
		var l, r int
		l, _ = strconv.Atoi(record[0])
		r, _ = strconv.Atoi(record[1])
		left = append(left, l)
		right = append(right, r)

		lineNum++
	}

	//----------------------------------------------------------------------
	// main - calculate
	//----------------------------------------------------------------------
	similarity := Similarity(left, right)
	fmt.Println("similarity:", similarity)
}

// Similarity returns the similarity of the two slices
// logic:
// 1. count the number(N) of elements same as the left each element(E[i]) in the right
// 2. score = Sum(N[i] * E[i])
func Similarity(left, right []int) int {
	score := 0
	for i := 0; i < len(left); i++ {
		// count the number of elements same as the left each element in the right
		count := CountNumberOfSameValues(left[i], right)
		score += (count * left[i])

		if DEBUG {
			fmt.Printf("left[%d]: %d, count: %d, score: %d\n", i, left[i], count, score)
		}
	}
	return score
}

func CountNumberOfSameValues(target int, slice []int) int {
	count := 0
	for i := 0; i < len(slice); i++ {
		if slice[i] == target {
			count++
		}
	}
	return count
}
