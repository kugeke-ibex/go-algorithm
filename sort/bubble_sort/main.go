package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	fmt.Println("Bubble Sort")
	arr := make([]int, 10)
	for i := 0; i < 10; i++ {
		arr[i] = rand.Intn(10)
	}
	start := time.Now()
	fmt.Println("Unsorted array:", arr)
	arr = bubbleSort(arr)
	end := time.Since(start)
	fmt.Println("Sorted array:", arr)
	fmt.Println(end)
}

func bubbleSort(arr []int) []int {
	for i := 0; i < len(arr) - 1; i++ {
		for j := 0; j < len(arr) - 1 - i; j++ {
			if arr[j] > arr[j+1] {
				arr[j], arr[j+1] = arr[j+1], arr[j]
				fmt.Println("Swapped array:", arr)
			}
		}
	}
	return arr
}
