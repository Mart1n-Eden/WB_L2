package main

import (
	"fmt"
	"strconv"
	"strings"
	"unicode"
)

func unpackString(s string) (string, error) {
	var (
		res   string
		char  byte
		count int
	)
	count = -1
	for i := 0; i < len(s); i++ {
		if unicode.IsDigit(rune(s[i])) {
			countStr := ""
			for unicode.IsDigit(rune(s[i])) {
				countStr += string(s[i])
				i++
				if i == len(s) {
					break
				}
			}
			i--
			count, _ = strconv.Atoi(countStr)
			res += strings.Repeat(string(char), count-1)
		} else {
			res += string(s[i])
			char = s[i]
		}
	}
	return res, nil
}

func main() {
	var s string
	fmt.Scan(&s)
	unpacked, err := unpackString(s)
	if err != nil {
		fmt.Printf("Ошибка: %v\n", err)
	} else {
		fmt.Println(unpacked)
	}

}
