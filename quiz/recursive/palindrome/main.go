package main

import (
	"fmt"
)

func main() {
	fmt.Printf("level is palindrome: %v\n", isPalindrome("level"))
	fmt.Printf("hello is palindrome: %v\n", isPalindrome("hello"))

	target := "fdafiewaafcabacdfafdaf"
	fmt.Printf("all palindromes of %s: %v\n", target, findAllPalindrome(target))
}

func isPalindrome(strings string) bool {
	length := len(strings)
	if length == 0 {
		return false
	}
	if length == 1 {
		return true
	}

	start, end := 0, length - 1
	if start < end {
		if strings[start] != strings[end] {
			return false
		}
		start++
		end--
	}
	return true
}

func findPalindrome(strings string, left int, right int) []string {
	result := make([]string, 0)
	for 0 <= left && right <= len(strings) - 1 {
		if strings[left] != strings[right] {
			break
		}
		result = append(result, strings[left: right+1])
		left--
		right++
	}

	return result
}

func findAllPalindrome(strings string) []string {
	result := make([]string, 0)
	length := len(strings)
	if length == 0 {
		return result
	}
	if length == 1 {
		result = append(result, strings)
	}

	for i := 1; i < length - 1; i++ {
		result = append(result, findPalindrome(strings, i-1,i+1)...)
		result = append(result, findPalindrome(strings, i-1, i)...)	
	}
	return result
}