package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	fmt.Println("Insertion Sort")
	arr := make([]int, 10)
	for i := 0; i < 10; i++ {
		arr[i] = rand.Intn(10)
	}

	start := time.Now()
	fmt.Println("Unsorted array:", arr)
	arr = insertionSort(arr)
	end := time.Since(start)
	fmt.Println("Sorted array:", arr)
	fmt.Println(end)
}

func insertionSort(arr []int) []int {
	lenArray := len(arr)
	for i := 1; i < lenArray; i++ {
		temp := arr[i]
		j := i - 1
		for j >= 0 && arr[j] > temp {
			arr[j+1] = arr[j]
			j--
			fmt.Println("Swapped array:", arr)
		}

		arr[j+1] = temp
		fmt.Println("Swapped temp array:", arr)
	}

	return arr
}