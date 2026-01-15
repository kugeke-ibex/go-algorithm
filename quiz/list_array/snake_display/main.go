package main


import (
	"fmt"
	"strings"
)

func main() {
	for _, row := range snakeStringV1("0123456789") {
		fmt.Println(strings.Join(row, ""))
	}


	fmt.Println("--------------------------------")	
	alphabetString := "abcdefghijklmnopqrstuvwxyz"
	for i := 0; i < 2; i++ {
		alphabetString = alphabetString + alphabetString
	}
	for _, row := range snakeStringV2(alphabetString, 10) {
		fmt.Println(strings.Join(row, ""))
	}
}

func snakeStringV1(chars string) [][]string {
	result := make([][]string, 3)
	resultIndexes := []int{0, 1, 2}
	insertIndex := 1

	for i, value := range chars {
		if i % 4 == 1 {
			insertIndex = 0
		} else if i % 2 == 0 {
			insertIndex = 1
		} else if i % 4 == 3 {
			insertIndex = 2
		}
		result[insertIndex] = append(result[insertIndex], string(value))

		for _, index := range resultIndexes {
			if index != insertIndex {
				result[index] = append(result[index], " ")
			}
		}
	}

	return result
}

func positive(a int) int {
	return a + 1
}

func negative(a int) int {
	return a - 1
}

func snakeStringV2(chars string, depth int) [][]string {
	result := make([][]string, depth)
	resultIndexes := make([]int, depth)
	for i := 0; i < depth; i++ {
		resultIndexes[i] = i
	}
	insertIndex := depth / 2

	op := negative
	for _, value := range chars {
		result[insertIndex] = append(result[insertIndex], string(value))
		for _, index := range resultIndexes {
			if index != insertIndex {
				result[index] = append(result[index], " ")
			}
		}
		
		if insertIndex <= 0 {
			op = positive
		} else if insertIndex >= depth -1 {
			op = negative
		}
		insertIndex = op(insertIndex)
	}

	return result
}
