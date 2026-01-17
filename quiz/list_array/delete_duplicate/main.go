package main

import (
	"fmt"
)

func main() {
	numbers := []int{1, 3, 3, 5, 5, 7, 7, 7, 10, 12, 12, 15}
	deleteDuplicateV1(&numbers)
	fmt.Println(numbers)
	numbers = []int{1, 3, 3, 5, 5, 7, 7, 7, 10, 12, 12, 15}
	deleteDuplicateV2(&numbers)
	fmt.Println(numbers)
	numbers = []int{1, 3, 3, 5, 5, 7, 7, 7, 10, 12, 12, 15}
	deleteDuplicateV3(&numbers)
	fmt.Println(numbers)
	numbers = []int{1, 3, 3, 5, 5, 7, 7, 7, 10, 12, 12, 15}
	deleteDuplicateV4(&numbers)
	fmt.Println(numbers)
}

func deleteDuplicateV1(numbers *[]int) {
	temp := make([]int, 0, len(*numbers))
	MapNumbers := make(map[int]bool)
	for _, number := range *numbers {
		if _, ok := MapNumbers[number]; !ok {
			MapNumbers[number] = true
			temp = append(temp, number)
		}
	}

	*numbers = temp
}

func deleteDuplicateV2(numbers *[]int) {
	temp := []int{(*numbers)[0]}
	i, lenNum := 0, len(*numbers) - 1
	for i < lenNum {
		if (*numbers)[i] != (*numbers)[i+1] {
			temp = append(temp, (*numbers)[i+1])
		}
		i++
	}

	*numbers = temp
}

func deleteDuplicateV3(numbers *[]int) {
	i := 0
	for i < len(*numbers) - 1 {
		if (*numbers)[i] == (*numbers)[i+1] {
			*numbers = append((*numbers)[:i], (*numbers)[i+1:]...)
			i--
		}
		i++
	}
}

func deleteDuplicateV4(numbers *[]int) {
	i := len(*numbers) - 1
	for i > 0 {
		if (*numbers)[i] == (*numbers)[i-1] {
			*numbers = append((*numbers)[:i], (*numbers)[i+1:]...)
		}
		i--
	}
}