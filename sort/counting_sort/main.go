package main

import (	
	"fmt"
	"math/rand"
	"time"
)

func main() {
	fmt.Println("Counting Sort")
	arr := make([]int, 10)
	for i := 0; i < 10; i++ {
		arr[i] = rand.Intn(10)
	}	

	start := time.Now()
	fmt.Println("Unsorted array:", arr)
	arr = countingSort(arr)
	end := time.Since(start)
	fmt.Println("Sorted array:", arr)
	fmt.Println(end)
}

func countingSort(arr []int) []int {
	maxValue := 0
	for i := 0; i < len(arr); i++ {
		if arr[i] > maxValue {
			maxValue = arr[i]
		}
	}
	fmt.Println("Max value:", maxValue)

	countArray := make([]int, maxValue + 1)
	for i := 0; i < len(arr); i++ {
		countArray[arr[i]]++
	}
	fmt.Println("Count array:", countArray)
	for i := 1; i < len(countArray); i++ {
		countArray[i] += countArray[i - 1]
	}
	fmt.Println("Count array after sum:", countArray)

	resultArray := make([]int, len(arr))
	for i := len(arr) - 1; i >= 0; i-- {
		index := countArray[arr[i]] - 1
		fmt.Println("Index:", index)
		resultArray[index] = arr[i]
		fmt.Println("Inserted array:", resultArray)
		countArray[arr[i]]--
	}

	return resultArray
}