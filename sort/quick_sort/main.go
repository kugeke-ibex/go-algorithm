package main


import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	fmt.Println("Quick Sort")
	arr := make([]int, 10)
	for i := 0; i < 10; i++ {
		arr[i] = rand.Intn(100)
	}

	start := time.Now()
	fmt.Println("Unsorted array:", arr)
	quickSort(arr)
	end := time.Since(start)
	fmt.Println("Sorted array:", arr)
	fmt.Println(end)
}

func partition(arr []int, low int, high int) int {
	i := low - 1
	pivot := arr[high]
	fmt.Println("Pivot:", pivot)

	for j := low; j < high; j++ {
		if arr[j] <= pivot {
			i++
			arr[i], arr[j] = arr[j], arr[i]
			fmt.Println("Swapped array:", arr)
		}
	}
	arr[i+1], arr[high] = arr[high], arr[i+1]
	fmt.Println("Swapped final array:", arr)

	return i + 1
}

func _quickSort(arr []int, low int, high int) {
	if low < high {
		partitionIndex := partition(arr, low, high)
		fmt.Println("Partition index:", partitionIndex)
		_quickSort(arr, low, partitionIndex - 1)
		_quickSort(arr, partitionIndex + 1, high)
	}

}

func quickSort(arr []int) []int {
	_quickSort(arr, 0, len(arr) - 1)

	return arr
}

