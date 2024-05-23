package main

import (
	"fmt"
	"sort"
	"strings"
)

func findAnagrams(words *[]string) map[string][]string {
	anagramSets := make(map[string][]string)

	for _, word := range *words {
		lowerWord := strings.ToLower(word)

		sortedWord := sortString(lowerWord)

		anagramSets[sortedWord] = append(anagramSets[sortedWord], lowerWord)
	}

	for key, value := range anagramSets {
		if len(value) <= 1 {
			delete(anagramSets, key)
		} else {
			sort.Strings(value)
			anagramSets[key] = value
		}
	}

	return anagramSets
}

func sortString(s string) string {
	runes := []rune(s)
	sort.Slice(runes, func(i, j int) bool {
		return runes[i] < runes[j]
	})
	return string(runes)
}

func main() {
	words := []string{"пятак", "пятка", "тяпка", "листок", "слиток", "столик", "мышь", "шум", "мыша"}

	anagramSets := findAnagrams(&words)

	for _, value := range anagramSets {
		fmt.Printf("%s: %v\n", value[0], value)
	}
}
