package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	fmt.Println("Radix Sort")
	arr := make([]int, 10)
	for i := 0; i < 10; i++ {
		arr[i] = rand.Intn(100)
	}	

	start := time.Now()
	fmt.Println("Unsorted array:", arr)
	arr = radixSort(arr)
	end := time.Since(start)
	fmt.Println("Sorted array:", arr)
	fmt.Println(end)
}

func radixSort(arr []int) []int {
	maxValue := 0;
	for i := 0; i < len(arr); i++ {
		if arr[i] > maxValue {
			maxValue = arr[i]
		}
	}
	fmt.Println("Max value:", maxValue)

	for digit := 1; maxValue / digit > 0; digit *= 10 {
		fmt.Println("Digit:", digit)
		arr = countingSort(arr, digit)
	}

	return arr
}

func countingSort(arr []int, digit int) []int {
	countArray := make([]int, 10)
	for i := 0; i < len(arr); i++ {
		index := (arr[i] / digit) % 10
		countArray[index]++
	}
	fmt.Println("Count array:", countArray)
	for i := 1; i < len(countArray); i++ {
		countArray[i] += countArray[i - 1]
	}
	fmt.Println("Count array after sum:", countArray)

	digitArray := make([]int, len(arr))
	for i := len(arr) - 1; i >= 0; i-- {
		index := (arr[i] / digit) % 10
		digitArray[countArray[index] - 1] = arr[i]
		countArray[index]--
	}
	fmt.Println("Digit array:", digitArray)


	return digitArray
}