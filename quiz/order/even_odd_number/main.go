package main

import (
	"fmt"
)

func main() {
	numbers := []int{0, 1, 3, 4, 2, 4, 5, 1, 6, 9, 8}
	orderEvenFirstOddLastV1(&numbers)
	fmt.Printf("ordered numbers: %v\n", numbers)
	numbers = []int{0, 1, 3, 4, 2, 4, 5, 1, 6, 9, 8}
	orderEvenFirstOddLastV2(&numbers)
	fmt.Printf("ordered numbers: %v\n", numbers)
}

func orderEvenFirstOddLastV1(numbers *[]int) {
	evenList, oddList := make([]int, 0), make([]int, 0)
	for _, number := range *numbers {
		if number % 2 == 0 {
			evenList = append(evenList, number)
		} else {
			oddList = append(oddList, number)
		}
	}

	*numbers = append(evenList, oddList...)
}

func orderEvenFirstOddLastV2(numbers *[]int) {
	i, j := 0, len(*numbers) - 1
	for i < j {
		if (*numbers)[i] % 2 == 0 {
			i++
		} else {
			(*numbers)[i], (*numbers)[j] = (*numbers)[j], (*numbers)[i]
			j--
		}
	}
}

