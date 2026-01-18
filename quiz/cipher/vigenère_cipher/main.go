package main

import (
	"fmt"
	"unicode"
	// "strings"
)

const alphabet = "ABCDEFGHIJKLMNOPQRSTUVWXYZ"

func main() {
	message := "ATTACK SILICON VALLEY"
	keyword := "HELLO"
	key := generateKey(message, keyword)
	fmt.Printf("Key: %s\n", key)
	fmt.Printf("Message: %s\n", message)
	encryptedMessage := encryptVigenere(message, key)
	fmt.Printf("Encrypted Message: %s\n", encryptedMessage)
	fmt.Printf("Decrypted Message: %s\n", decryptVigenere(encryptedMessage, key))
}

func generateKey(message string, keyword string) string {
	key := keyword
	remainLength := len(message) - len(keyword)
	for i := range remainLength {
		key += string(keyword[i % len(keyword)])
	}

	return key
}

func encryptVigenere(message string, key string) string {
	result := ""
	for i, char := range message {
		if !unicode.IsLetter(char) {
			result += string(char)
			continue
		}

		// index := (strings.Index(alphabet, string(char)) + strings.Index(alphabet, string(key[i]))) % len(alphabet)
		index := (int(unicode.ToUpper(char)) + int(unicode.ToUpper(rune(key[i])))) % len(alphabet)
		result += string(alphabet[index])
	}

	return result
}

func decryptVigenere(message string, key string) string {
	result := ""
	for i, char := range message {
		if !unicode.IsLetter(char) {
			result += string(char)
			continue
		}

		// index := ((strings.Index(alphabet, string(char))) - (strings.Index(alphabet, string(key[i]))) + len(alphabet)) % len(alphabet)
	
		index := (int(unicode.ToUpper(char)) - int(unicode.ToUpper(rune(key[i]))) + len(alphabet)) % len(alphabet)
		result += string(alphabet[index])
	}

	return result
}

