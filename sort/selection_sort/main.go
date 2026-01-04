package main


import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	fmt.Println("Selection Sort")
	arr := make([]int, 10)
	for i := 0; i < 10; i++ {
		arr[i] = rand.Intn(10)
	}

	start := time.Now()
	fmt.Println("Unsorted array:", arr)
	arr = selectionSort(arr)
	end := time.Since(start)
	fmt.Println("Sorted array:", arr)
	fmt.Println(end)
}

func selectionSort(arr []int) []int {
	lenArray := len(arr)
	for i := 0; i < lenArray - 1; i++ {
		minIndex := i
		for j := i + 1; j < lenArray - 1; j++ {
			if arr[j] < arr[minIndex] {
				minIndex = j
			}
		}

		fmt.Println("Min index:", minIndex)
		arr[i], arr[minIndex] = arr[minIndex], arr[i]
		fmt.Println("Swapped array:", arr)
	}

	return arr
}

