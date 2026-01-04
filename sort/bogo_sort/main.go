package main

import (
	"math/rand"
	"fmt"
	"time"
)

func main() {
	fmt.Println("Bogosort")
	arr := make([]int, 10)
	for i := 0; i < 10; i++ {
		arr[i] = rand.Intn(10)
	}
	fmt.Println(arr)
	start := time.Now()
	arr = bogoSort(arr)
	end := time.Since(start)
	fmt.Println(arr)
	fmt.Println(end)
}

func bogoSort(arr []int) []int {
	for !isSorted(arr) {
		rand.Shuffle(len(arr), func(i, j int) {
			arr[i], arr[j] = arr[j], arr[i]
		})
	}
	return arr
}

func isSorted(arr []int) bool {
	for i := 0; i < len(arr)-1; i++ {
		if arr[i] > arr[i+1] {
			return false
		}
	}
	return true
}