package main

import (
	"fmt"
)

func main() {
	fmt.Println(getMaxSequenceSum([]int{1, -2, 3, 6, -1, 2, 4, -6, 2}))
	fmt.Println(findMaxCircularSequenceSum([]int{1, -2, 3, 6, -1, 2, 4, -5, 2}))
}

func getMaxSequenceSum(numbers []int) int {
	resultSequence, sumSequence := 0, 0
	
	for _, number := range numbers {
		tempSumSequence := sumSequence + number
		if number < tempSumSequence {
			sumSequence = tempSumSequence
		} else {
			sumSequence = number
		}

		if resultSequence < sumSequence {
			resultSequence = sumSequence
		}
	}

	return resultSequence
}

func findMaxCircularSequenceSum(numbers []int) int {
	maxSequenceSum := getMaxSequenceSum(numbers)

	invertedNumbers := make([]int, len(numbers))
	allSum := 0
	for _, number := range numbers {
		allSum += number
		invertedNumbers = append(invertedNumbers, -number)
	}

	maxWrapSequence := allSum - (-getMaxSequenceSum(invertedNumbers))

	if maxWrapSequence > maxSequenceSum {
		return maxWrapSequence
	}

	return maxSequenceSum
}

