package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

func readLines(filename string) ([]string, error) {
	file, err := os.Open(filename)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	var lines []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}
	return lines, scanner.Err()
}

func sortByColumn(lines []string, column int, numeric bool, reverse bool, unique bool) []string {
	sort.SliceStable(lines, func(i, j int) bool {
		var a, b string
		if len(strings.Fields(lines[i])) > column {
			a = strings.Fields(lines[i])[column]
		}
		if len(strings.Fields(lines[j])) > column {
			b = strings.Fields(lines[j])[column]
		}
		if numeric {
			numa, erra := strconv.ParseFloat(a, 64)
			numb, errb := strconv.ParseFloat(b, 64)
			if erra == nil && errb == nil {
				if reverse {
					return numa > numb
				}
				return numa < numb
			}
		}
		if reverse {
			return a > b
		}
		return a < b
	})

	if unique {
		lines = uniqueLines(lines)
	}

	return lines
}

func uniqueLines(lines []string) []string {
	uniqueMap := make(map[string]bool)
	var uniqueLines []string
	for _, line := range lines {
		if !uniqueMap[line] {
			uniqueMap[line] = true
			uniqueLines = append(uniqueLines, line)
		}
	}
	return uniqueLines
}

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Usage: go run main.go <filename> [-k column] [-n] [-r] [-u]")
		os.Exit(1)
	}

	filename := os.Args[1]

	var column int
	numeric := false
	reverse := false
	unique := false

	for i := 2; i < len(os.Args); i++ {
		if os.Args[i] == "-k" && i+1 < len(os.Args) {
			col, err := strconv.Atoi(os.Args[i+1])
			if err != nil {
				fmt.Println("Invalid column number:", os.Args[i+1])
				os.Exit(1)
			}
			column = col
			i++
		} else if os.Args[i] == "-n" {
			numeric = true
		} else if os.Args[i] == "-r" {
			reverse = true
		} else if os.Args[i] == "-u" {
			unique = true
		}
	}

	lines, err := readLines(filename)
	if err != nil {
		fmt.Println("Error reading file:", err)
		os.Exit(1)
	}

	lines = sortByColumn(lines, column, numeric, reverse, unique)

	for _, line := range lines {
		fmt.Println(line)
	}
}
