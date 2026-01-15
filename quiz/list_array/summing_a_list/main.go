package main

import (
	"fmt"
	"math"
)

func main() {

	fmt.Printf("[1,2,3] -> %d\n", listToIntPlusOne([]int{1, 2, 3}))
	fmt.Printf("[1,2,9] -> %d\n", listToIntPlusOne([]int{1, 2, 9}))
	fmt.Printf("[1,9,9] -> %d\n", listToIntPlusOne([]int{1, 9, 9}))
	fmt.Printf("[9,9,9] -> %d\n", listToIntPlusOne([]int{9, 9, 9}))
}

func removeZero(numbers *[]int) {
	if len(*numbers) != 0 && (*numbers)[0] == 0 {
		*numbers = (*numbers)[1:]
		removeZero(numbers)
	}
}

func listToInt(numbers []int) int {
	result := 0
	for i := len(numbers) - 1; i >= 0; i-- {
		result += numbers[i] * int(math.Pow(10, float64(len(numbers) - i - 1)))
	}

	return result
}

func listToIntPlusOne(numbers []int) int {
	i := len(numbers) - 1
	numbers[i] += 1

	for 0 < i {
		if numbers[i] != 10 {
			removeZero(&numbers)
			break
		}

		numbers[i] = 0
		numbers[i-1] += 1
		i -= 1
	}

	if numbers[0] == 10 {
		numbers[0] = 1
		numbers = append(numbers, 0)
	}

	return listToInt(numbers)
}