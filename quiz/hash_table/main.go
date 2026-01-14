package main

import (
	"fmt"
)

func main() {
	x := []int{1, 2, 3, 4, 4, 5, 5, 8, 10}
	y := []int{4, 5, 5, 5, 6, 7, 8, 8, 10}
	fmt.Println(x, y)
	minCountRemove(&x, &y)
	fmt.Println(x, y)
}

func minCountRemove(x *[]int, y *[]int) {
	countX := make(map[int]int)
	countY := make(map[int]int)

	for _, value := range *x {
		countX[value]++
	}
	for _, value := range *y {
		countY[value]++
	}
	
	for keyx, valuex := range countX {
		// fmt.Println(keyx, valuex)
		if valuey, ok := countY[keyx]; ok {
			if valuex > valuey {
				*y = deleteElement(*y, keyx)
			} else if valuex < valuey {
				*x = deleteElement(*x, keyx)
			}
		}
	}
}

func deleteElement(slice []int, target int) []int {

	newSlice := make([]int, 0, len(slice))
	for _, value := range slice {
		if value == target {
			continue
		}

		newSlice = append(newSlice, value)
	}
	return newSlice
}