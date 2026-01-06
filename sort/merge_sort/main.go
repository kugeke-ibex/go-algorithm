package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	fmt.Println("Merge Sort")
	arr := make([]int, 10)
	for i := 0; i < 10; i++ {
		arr[i] = rand.Intn(100)
	}

	start := time.Now()
	fmt.Println("Unsorted array:", arr)
	arr = mergeSort(arr)
	end := time.Since(start)
	fmt.Println("Sorted array:", arr)
	fmt.Println(end)
}

func mergeSort(arr []int) []int {
	if len(arr) <= 1 {
		return arr
	}

	center := len(arr) / 2
	left := arr[:center]
	right := arr[center:]

	fmt.Println("center:", center)
	fmt.Println("Left array:", left)
	fmt.Println("Right array:", right)
	left = mergeSort(left)
	right = mergeSort(right)

	i := 0
	j := 0
	k := 0

	resultArray := make([]int, len(arr))
	for i < len(left) && j < len(right) {
		if left[i] <= right[j] {
			resultArray[k] = left[i]
			i++
		} else {
			resultArray[k] = right[j]
			j++
		}
		k++
	}

	for i < len(left) {
		resultArray[k] = left[i]
		i++
		k++
	}

	for j < len(right) {
		resultArray[k] = right[j]
		j++
		k++
	}
	fmt.Println("Result array:", resultArray)

	return resultArray
}
