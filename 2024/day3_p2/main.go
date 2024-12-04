package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"
)

var DEBUG = true

func main() {
	// read data from stdin
	data := readStdinAsOneLine()

	// parse data - get mul(a, b) as a slice
	calls := getMultipleCallsWithCond(data)
	if calls == nil || len(calls) == 0 {
		fmt.Println("no mul(a, b) calls found")
		return
	}

	if DEBUG {
		fmt.Printf("calls: %v\n", calls)
	}

	// calculate
	total := 0
	do := true
	for _, call := range calls {
		curExec := ""
		curResult := 0

		// parse the call
		mul := mulString(call)
		if mul {
			// e.g. "mul(2,3) under do=true" -> 6
			// e.g. "mul(2,3) under do=false" -> skip
			if do {
				result, err := mulFromString(call)
				if err != nil {
					fmt.Printf("error: %v\n", err)
					return
				}
				total += result

				curExec = "mul"
				curResult = result
			} else {
				// skip
				curExec = "skip"
			}
		} else {
			// do() or don't()
			// e.g. "do()" -> do = true
			// e.g. "don't()" -> do = false
			do = enableDoFromString(call)

			curExec = call
		}

		if DEBUG {
			fmt.Printf("call: %s, result: %d, exec: %s\n", call, curResult, curExec)
		}
	}
	fmt.Println("total:", total)
}

func readStdinAsOneLine() string {
	scanner := bufio.NewScanner(os.Stdin)
	lines := make([]string, 0)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}

	// join all lines
	// use strings.Builder for memory efficiency
	b := strings.Builder{}
	for _, line := range lines {
		b.WriteString(line)
	}
	return b.String()
}

func getMultipleCalls(data string) []string {
	// get all mul(a, b) calls
	// use regexp package
	exp := `mul\(\d+,\d+\)`
	re := regexp.MustCompile(exp)

	return re.FindAllString(data, -1)
}

func getMultipleCallsWithCond(data string) []string {
	// get all mul(a, b) calls with conditional statements (do(), don't())
	// use regexp package
	exp := `mul\(\d+,\d+\)|do\(\)|don't\(\)`
	re := regexp.MustCompile(exp)

	return re.FindAllString(data, -1)
}

func mul(a, b int) int {
	return a * b
}

// mulFromString parses the string and returns the result of the multiplication
// e.g. mulFromString("mul(2,3)") returns 6
func mulFromString(s string) (int, error) {
	exp := `^mul\((\d+),(\d+)\)$`
	re := regexp.MustCompile(exp)

	// maches like ["mul(2,3)", "2", "3"]
	matches := re.FindStringSubmatch(s)
	if matches == nil || len(matches) != 3 {
		return 0, fmt.Errorf("invalid input string: %s, matches: %v", s, matches)
	}
	a, err1 := strconv.Atoi(matches[1])
	b, err2 := strconv.Atoi(matches[2])
	if err1 != nil || err2 != nil {
		return 0, fmt.Errorf("invalid input string: %s (err1: %v, err2: %v)", s, err1, err2)
	}

	return mul(a, b), nil
}

func mulString(s string) bool {
	return strings.Contains(s, "mul")
}

func enableDoFromString(s string) bool {
	return strings.Contains(s, "do()")
}
