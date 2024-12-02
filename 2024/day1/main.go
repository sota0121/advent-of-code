package main

import (
	"bufio"
	"encoding/csv"
	"fmt"
	"math"
	"os"
	"slices"
	"strconv"
)

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
	slices.Sort(left)
	slices.Sort(right)

	absSum := 0
	// for high performance, we use the index instead of the value
	for i := 0; i < len(left); i++ {
		absSum += int(math.Abs(float64(left[i]) - float64(right[i])))
	}
	fmt.Println("Abs sum:", absSum)
}
