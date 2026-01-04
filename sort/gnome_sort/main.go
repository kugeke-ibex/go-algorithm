package main


import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	fmt.Println("Gnome Sort")
	arr := make([]int, 10)
	for i := 0; i < 10; i++ {
		arr[i] = rand.Intn(10)
	}

	start := time.Now()
	fmt.Println("Unsorted array:", arr)
	arr = gnomeSort(arr)
	end := time.Since(start)
	fmt.Println("Sorted array:", arr)
	fmt.Println(end)
}

func gnomeSort(arr []int) []int {
	lenArray := len(arr)
	index := 0
	for index < lenArray {
		if index == 0 {
			index++
		}
		if arr[index] >= arr[index-1] {
			index++
		} else {
			arr[index], arr[index-1] = arr[index-1], arr[index]
			index--
			fmt.Println("Swapped array:", arr)
		}
	}

	return arr
}

