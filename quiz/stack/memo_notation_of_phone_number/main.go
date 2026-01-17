package main

import (
	"fmt"
	"strings"
)

type PhoneAlphabet string

var phoneNumberMap = map[string]string {
	"0": "+",
	"1": "@",
	"2": "ABC",
	"3": "DEF",
	"4": "GHI",
	"5": "JKL",
	"6": "MNO",
	"7": "PQRS",
	"8": "TUV",
	"9": "WXYZ",
} 

func main() {
	fmt.Printf("phonemnemonic: %v\n", phoneMnemonicV1("23"))

	for _, s := range phoneMnemonicV1("568-379-8466") {
		if s == "LOVEPYTHON" {
			fmt.Printf("phoneMnemonicV1: %s\n", s)
		} 
	}

	for _, s := range phoneMnemonicV2("568-379-8466") {
		if s == "LOVEPYTHON" {
			fmt.Printf("phoneMnemonicV2: %s\n", s)
		} 
	}

}


func findCandidateAlphabet(digit int, phoneNumberList []string, candidate *[]string, temp *[]string) {
	if digit == len(phoneNumberList) {
		// fmt.Printf("digit: %d, phoneNumber: %s, candidate: %v, temp: %v\n", digit, phoneNumberList, *candidate, *temp)
		*candidate = append((*candidate), strings.Join((*temp), ""))
	} else {
		// fmt.Printf("phoneNumberMap[rune(phoneNumber[digit])]: %s\n", phoneNumberMap[phoneNumberList[digit]])
		for _, char := range phoneNumberMap[phoneNumberList[digit]] {
			(*temp)[digit] = string(char)
			// fmt.Printf("temp: %v\n", (*temp))
			findCandidateAlphabet(digit + 1, phoneNumberList, candidate, temp)
		}
	}
}

func phoneMnemonicV1(phoneNumber string) []string {
	phoneNumber = strings.ReplaceAll(phoneNumber, "-", "")
	phoneNumberList := make([]string, len(phoneNumber))
	for i, char := range phoneNumber {
		phoneNumberList[i] = string(char)
	}
	candidate := make([]string, 0)
	temp := make([]string, len(phoneNumber))
	findCandidateAlphabet(0, phoneNumberList, &candidate, &temp)
	return candidate
}

func phoneMnemonicV2(phoneNumber string) []string {
	phoneNumber = strings.ReplaceAll(phoneNumber, "-", "")
	phoneNumberList := make([]string, len(phoneNumber))
	for i, char := range phoneNumber {
		phoneNumberList[i] = string(char)
	}
	candidate := make([]string, 0)
	stack := make([]string, 0)
	stack = append(stack, "")

	for len(stack) != 0 {
		alphabets := stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		if len(alphabets) == len(phoneNumberList) {
			candidate = append(candidate, alphabets)
		} else {
			for _, char := range phoneNumberMap[phoneNumberList[len(alphabets)]] {
				stack = append(stack, alphabets + string(char))
			}
		}
	}

	return candidate
}
	
