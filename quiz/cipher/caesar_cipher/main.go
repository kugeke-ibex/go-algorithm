package main

import (
	"fmt"
	"strings"
	"unicode"
)

func main() {
	rawString := "ATTACK SILICON VALLEY by enginner"
	fmt.Printf("%s -> caesarCipherV1 shift 3: %s\n", rawString, caesarCipherV1(rawString, 3))
	fmt.Printf("%s -> caesarCipherV2 shift 3: %s\n", rawString, caesarCipherV2(rawString, 3))
	fmt.Println()
	encryptedString := caesarCipherV1(rawString, 3)
	fmt.Printf("%s -> caesarCipherV1 shift -3: %s\n", encryptedString, caesarCipherV1(encryptedString, -3))
	fmt.Printf("%s -> caesarCipherV2 shift -3: %s\n", encryptedString, caesarCipherV2(encryptedString, -3))

	fmt.Println()
	fmt.Println("caesarCipherHackV1:")
	for _, hack := range caesarCipherHackV1(encryptedString) {
		for shift, text := range hack {
			fmt.Printf("Shift: %d, Text: %s\n", shift, text)
		}
	}
	fmt.Println()
	fmt.Println("caesarCipherHackV2:")
	for _, hack := range caesarCipherHackV2(encryptedString) {
		for shift, text := range hack {
			fmt.Printf("Shift: %d, Text: %s\n", shift, text)
		}
	}
}


func caesarCipherV1(text string, shift int) string {
	result := ""
	var alphabet string
	for _, char := range text {
		if unicode.IsUpper(char) {
			alphabet = "ABCDEFGHIJKLMNOPQRSTUVWXYZ"
		} else if unicode.IsLower(char) {
			alphabet = "abcdefghijklmnopqrstuvwxyz"
		} else {
			result += string(char)
			continue
		}
		if shift < 0 {
			shift = len(alphabet) + shift
		}

		index := (strings.Index(alphabet, string(char)) + shift) % len(alphabet)
		result += string(alphabet[index])
 	}

	return result
}

func caesarCipherV2(text string, shift int) string {
	result := ""
    lengthAlphabet := 'Z' - 'A' + 1
	if shift < 0 {
		shift = int(lengthAlphabet) + shift
	}
	
	for _, char := range text {
		if unicode.IsUpper(char) {
			result += string((char + rune(shift) - 'A') % lengthAlphabet + 'A')
		} else if unicode.IsLower(char) {
			result += string((char + rune(shift) - 'a') % lengthAlphabet + 'a')
		} else {
			result += string(char)
		}
 	}

	return result
}

func caesarCipherHackV1(text string) []map[int]string {
	result := make([]map[int]string, 0)
	lengthAlphabet := int('Z' - 'A' + 1)

	var alphabet string
	for shift := 1; shift <= lengthAlphabet; shift++ {
		revertedText := ""
		for _, char := range text {
			if unicode.IsUpper(char) {
				alphabet = "ABCDEFGHIJKLMNOPQRSTUVWXYZ"
			} else if unicode.IsLower(char) {
				alphabet = "abcdefghijklmnopqrstuvwxyz"
			} else {
				revertedText += string(char)
				continue
			}

			index := strings.Index(alphabet, string(char)) - shift
			if index < 0 {
				index += lengthAlphabet
			}

			revertedText += string(alphabet[index])
		}

		result = append(result, map[int]string{shift: revertedText})
	}

	return result
}

func caesarCipherHackV2(text string) []map[int]string {
	result := make([]map[int]string, 0)
	lengthAlphabet := 'Z' - 'A' + 1

	for shift := 1; shift <= int(lengthAlphabet); shift++ {
		revertedText := ""
		for _, char := range text {
			if unicode.IsUpper(char) {
				index := char - rune(shift)
				if index < 'A' {
					index += lengthAlphabet
				}
				revertedText += string(index)
			} else if unicode.IsLower(char) {
				index := char - rune(shift)
				if index < 'a' {
					index += lengthAlphabet
				}
				revertedText += string(index)
			} else {
				revertedText += string(char)
			}
		}

		result = append(result, map[int]string{shift: revertedText})
	}

	return result
}