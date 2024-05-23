package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"
	"regexp"
	// "strings"
)

func main() {
	// Парсинг аргументов командной строки
	// after := flag.Int("A", 0, "Print N lines after the match")
	// before := flag.Int("B", 0, "Print N lines before the match")
	context := flag.Int("C", 0, "Print ±N lines around the match")
	count := flag.Bool("c", false, "Count the number of matching lines")
	ignoreCase := flag.Bool("i", false, "Ignore case when matching")
	invert := flag.Bool("v", false, "Invert the match")
	fixed := flag.Bool("F", false, "Fixed string match (no pattern)")
	lineNum := flag.Bool("n", false, "Print line numbers")
	flag.Parse()

	// Получение паттерна для поиска
	pattern := flag.Arg(0)

	// Создание регулярного выражения в зависимости от флагов
	var regex *regexp.Regexp
	if *fixed {
		pattern = regexp.QuoteMeta(pattern)
	}
	if *ignoreCase {
		regex = regexp.MustCompile("(?i)" + pattern)
	} else {
		regex = regexp.MustCompile(pattern)
	}

	// Открытие файла для поиска или использование стандартного ввода
	var file *os.File
	if flag.NArg() > 1 {
		var err error
		file, err = os.Open(flag.Arg(1))
		if err != nil {
			fmt.Println("Error opening file:", err)
			os.Exit(1)
		}
		defer file.Close()
	} else {
		file = os.Stdin
	}

	// Сканирование файла построчно и применение фильтра
	scanner := bufio.NewScanner(file)
	var matchingLines int
	var outputLines []string
	for lineNumber := 1; scanner.Scan(); lineNumber++ {
		line := scanner.Text()
		matched := regex.MatchString(line)
		if (*invert && !matched) || (!*invert && matched) {
			if *count {
				matchingLines++
			} else {
				if *lineNum {
					line = fmt.Sprintf("%d:%s", lineNumber, line)
				}
				outputLines = append(outputLines, line)
				if *context > 0 {
					outputLines = append(outputLines, getContextLines(scanner, *context)...)
				}
			}
		}
	}

	// Вывод результата
	if *count {
		fmt.Println(matchingLines)
	} else {
		for _, line := range outputLines {
			fmt.Println(line)
		}
	}
}

// Получить N строк контекста вокруг совпадения
func getContextLines(scanner *bufio.Scanner, context int) []string {
	var lines []string
	for i := 0; i < context && scanner.Scan(); i++ {
		lines = append(lines, scanner.Text())
	}
	return lines
}
