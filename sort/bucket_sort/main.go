package main

import (	
	"fmt"
	"math/rand"
	"time"
)

func main() {
	fmt.Println("Bucket Sort")
	arr := make([]int, 10)
	for i := 0; i < 10; i++ {
		arr[i] = rand.Intn(10)
	}

	start := time.Now()
	fmt.Println("Unsorted array:", arr)
	arr = bucketSort(arr)
	end := time.Since(start)
	fmt.Println("Sorted array:", arr)
	fmt.Println(end)
}

func bucketSort(arr []int) []int {
	maxValue := 0
	for i := 0; i < len(arr); i++ {
		if arr[i] > maxValue {
			maxValue = arr[i]
		}
	}

	lenArray := len(arr)
	size := maxValue / lenArray
	fmt.Println("Max value:", maxValue)
	fmt.Println("Length of array:", lenArray)
	if size == 0 {
		size = 1
	}
	fmt.Println("Size:", size)
	buckets := make([][]int, size)
	for i := 0; i < lenArray; i++ {
		// 整数除算
		if size == 1 {
			buckets[0] = append(buckets[0], arr[i])
		} else {
			index := arr[i] / size
			if index != size {
				buckets[index] = append(buckets[index], arr[i])
			} else {
				buckets[size - 1] = append(buckets[size - 1], arr[i])
			}
		}
	}
	for i := 0; i < size; i++ {
		buckets[i] = insertionSort(buckets[i])
	}

	resultArray := make([]int, 0, lenArray)
	for i := 0; i < size; i++ {
		resultArray = append(resultArray, buckets[i]...)
	}

	return resultArray;
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