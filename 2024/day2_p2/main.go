package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

var DEBUG = false

func main() {
	// read data from stdin (space separated)
	// Not all reports include the same number of levels
	// Tha's why we can't use csv.NewReader, which assumes all records have the same number of fields
	// Instead, we use bufio.NewScanner and split the input data by space
	scanner := bufio.NewScanner(os.Stdin)

	numOfSafeReports := 0
	scanIndex := 0
	for scanner.Scan() {
		line := strings.Trim(scanner.Text(), " ")
		fields := strings.Fields(line)

		levels := make([]int, 0)
		for i := 0; i < len(fields); i++ {
			level, _ := strconv.Atoi(fields[i])
			levels = append(levels, level)
		}

		// If at least one report is safe, the system will be considered safe
		allLevels := generateWithAugmentedLists(levels)
		safe := false
		for i, levels := range allLevels {
			safe = IsSafeReport(levels)

			if DEBUG {
				fmt.Printf("scan: %d, augmented[%d]: %v, safe?: %v\n", scanIndex, i, levels, safe)
			}

			if safe {
				numOfSafeReports++
				break
			}
		}
		// if IsSafeReport(levels) {
		// 	numOfSafeReports++
		// }

		if DEBUG {
			fmt.Printf("scan: %d, levels: %v, safe?: %v\n", scanIndex, levels, safe)
		}

		scanIndex++
	}

	if err := scanner.Err(); err != nil {
		fmt.Println("Error:", err)
		panic(err)
	}
	fmt.Println("Number of safe reports:", numOfSafeReports)
}

// generateWithAugmentedLists generates all possible reports
// 1. original
// 2. for each level: remove one element
// e.g.
// original: [1, 2, 3]
// augmented: [[1, 2, 3], [2, 3], [1, 3], [1, 2]]
func generateWithAugmentedLists(levels []int) [][]int {
	// 本当は、要素数が最低でも３ないといけないというエラーチェックいれたいけど、今回は省略

	// if DEBUG {
	// 	fmt.Println("original:", levels)
	// }

	lists := make([][]int, 0, len(levels)+1)
	lists = append(lists, levels) // original

	for i := range levels {
		augmented := make([]int, 0, len(levels)-1)
		augmented = append(augmented, levels[:i]...)
		augmented = append(augmented, levels[i+1:]...)
		lists = append(lists, augmented)

		// if DEBUG {
		// 	fmt.Printf("augmented[%d]: %v\n", i, augmented)
		// }
	}

	return lists
}

// IsSafeReport returns true if the report is safe
// if the both conditions are true, the report is safe
// 1. all increasing or all decreasing (= Monotonic)
// 2. any 2 adjacent differ at least 1, at most 3
func IsSafeReport(levels []int) bool {
	return IsMonotonic(levels) && AdjacentDifference(levels)
}

func IsMonotonic(levels []int) bool {
	// all increasing or all decreasing
	isIncreasing := true
	isDecreasing := true

	for i := 1; i < len(levels); i++ {
		if levels[i-1] < levels[i] {
			isDecreasing = false
		}
		if levels[i-1] > levels[i] {
			isIncreasing = false
		}
		if !isIncreasing && !isDecreasing {
			return false
		}
	}
	return true
}

func AdjacentDifference(levels []int) bool {
	// any 2 adjacent differ at least 1, at most 3
	for i := 1; i < len(levels); i++ {
		if !IsInRange(levels[i-1], levels[i]) {
			return false
		}
	}
	return true
}

func IsInRange(a, b int) bool {
	diff := math.Abs(float64(a - b))
	// 1 <= diff <= 3
	if diff < 1 || diff > 3 {
		return false
	}
	return true
}
