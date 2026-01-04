package main


import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	fmt.Println("Cocktail Sort")
	arr := make([]int, 10)
	for i := 0; i < 10; i++ {
		arr[i] = rand.Intn(10)
	}

	start := time.Now()
	fmt.Println("Unsorted array:", arr)
	arr = cocktailSort(arr)
	end := time.Since(start)
	fmt.Println("Sorted array:", arr)
	fmt.Println(end)
}

func cocktailSort(arr []int) []int {
	lenArray := len(arr)
	swapped := true
	start := 0
	end := lenArray - 1
	for swapped {
		swapped = false
		// Forward: インデックス番号が小さい方から大きい方へ移動
		for i := start; i < end; i++ {
			if arr[i] > arr[i+1] {
				arr[i], arr[i+1] = arr[i+1], arr[i]
				swapped = true
				fmt.Println("Swapped array forward:", arr)
			}
		}

		if !swapped {
			break
		}
		end--

		// Backward: インデックス番号が大きい方から小さい方へ移動
		for i := end - 1; i >= start; i-- {
			if arr[i] > arr[i+1] {
				arr[i], arr[i+1] = arr[i+1], arr[i]
				swapped = true
				fmt.Println("Swapped array backward:", arr)
			}
		}

		if !swapped {
			break
		}
		start++
	}
	return arr
}