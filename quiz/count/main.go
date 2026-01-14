package main

import (
	"fmt"
	"strings"
	"unicode"
)

func main() {
	fmt.Println(countCharsV1("This is a pen. This is an apple. Applepen."))
	fmt.Println(countCharV2("This is a pen. This is an apple. Applepen."))
}

func countCharsV1(str string) map[string]int {
	charCount := make(map[string]int)
	for _, char := range strings.ToLower(str) {
		if unicode.IsSpace(char) {
			continue
		}
		charCount[string(char)]++
	}

	maxChar := make(map[string]int)
	maxCount := 0
	for char, count := range charCount {
		if count > maxCount {
			maxChar = map[string]int{char: count}
			maxCount = count
		}
	}
	return maxChar
}

func countCharV2(str string) map[string]int {
	charCount := make(map[string]int)
	maxChar := ""
	maxCount := 0

	for _, char := range strings.ToLower(str) {
		if unicode.IsSpace(char) {
			continue
		}
		charStr := string(char)
		charCount[charStr]++
		if charCount[charStr] > maxCount {
			maxCount = charCount[charStr]
			maxChar = charStr
		}
	}

	return map[string]int{maxChar: maxCount}
}
