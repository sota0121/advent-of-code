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
	calls := getMultipleCalls(data)
	if calls == nil || len(calls) == 0 {
		fmt.Println("no mul(a, b) calls found")
		return
	}

	if DEBUG {
		fmt.Printf("calls: %v\n", calls)
	}

	// calculate
	total := 0
	for _, call := range calls {
		// parse the call
		// e.g. "mul(2,3)" -> 6
		result, err := mulFromString(call)
		if err != nil {
			fmt.Printf("error: %v\n", err)
			return
		}
		total += result

		if DEBUG {
			fmt.Printf("call: %s, result: %d\n", call, result)
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
