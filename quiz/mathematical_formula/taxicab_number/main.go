package main 

import (
	"fmt"
	"math"
)

func main() {
	fmt.Printf("Taxicab Number(1, 2): %v\n", taxiCabNumber(1, 2, 1))
	fmt.Printf("Taxicab Number(2, 2): %v\n", taxiCabNumber(2, 2, 1))
	fmt.Printf("Taxicab Number(1, 3): %v\n", taxiCabNumber(1, 3, 87539319))
}

func taxiCabNumber(maxAnswerNumber int, matchAnswerNumber int, startAnswer int) []map[int][][]int {
	result := make([]map[int][][]int, 0)
	gotAnswerCount := 0
	answer := startAnswer
	if startAnswer <= 0 {
		answer = 1
	}

	for gotAnswerCount < maxAnswerNumber {
		memo := make(map[int][][]int)
		matchAnswerCount := 0

		maxParam := int(math.Pow(float64(answer), 1.0/3.0)) + 1
		for x := 1; x < maxParam; x++ {
			for y := x+1; y < maxParam; y++ {
				if math.Pow(float64(x), 3) + math.Pow(float64(y), 3) == float64(answer) {
					matchAnswerCount++
					memo[answer] = append(memo[answer], []int{x, y})
				}
			}
		}

		if matchAnswerCount == matchAnswerNumber {
			result = append(result, memo)
			gotAnswerCount++
		}
		answer++
	}
	return result
}
