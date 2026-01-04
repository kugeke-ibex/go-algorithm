package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	fmt.Println("Comb Sort")
	arr := make([]int, 10)
	for i := 0; i < 10; i++ {
		arr[i] = rand.Intn(10)
	}

	start := time.Now()
	fmt.Println("Unsorted array:", arr)
	arr = combSort(arr)
	end := time.Since(start)
	fmt.Println("Sorted array:", arr)
	fmt.Println(end)
}

func combSort(arr []int) []int {
	lenArray := len(arr)
	gap := lenArray
	swapped := true

	for gap != 1 || swapped {
		gap = int(float64(gap) / 1.3)
		fmt.Println("Gap:", gap)
		if gap < 1 {
			gap = 1
		}
		swapped = false

		for i := 0; i < lenArray - gap; i++ {
			if arr[i] > arr[i+gap] {
				arr[i], arr[i+gap] = arr[i+gap], arr[i]
				swapped = true
				fmt.Println("Swapped array:", arr)
			}
		}
	}

	return arr
}
