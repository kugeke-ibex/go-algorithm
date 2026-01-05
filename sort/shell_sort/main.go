package main


import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	fmt.Println("Shell Sort")
	arr := make([]int, 10)
	for i := 0; i < 10; i++ {
		arr[i] = rand.Intn(10)
	}	

	start := time.Now()
	fmt.Println("Unsorted array:", arr)
	shellSort(arr)
	end := time.Since(start)
	fmt.Println("Sorted array:", arr)
	fmt.Println(end)

}

func shellSort(arr []int) []int {
	lenArray := len(arr)
	
	for gap := lenArray / 2; gap > 0; {
		fmt.Println("Gap:", gap)
		for i := gap; i < lenArray; i++ {
			temp := arr[i]
			j := i
			for j >= gap && arr[j-gap] > temp {
				arr[j] = arr[j-gap]
				j -= gap
				fmt.Println("Swapped array:", arr)
			}
			arr[j] = temp 
			fmt.Println("Swapped temp array:", arr)
		}
		gap /= 2
	}

	return arr
}