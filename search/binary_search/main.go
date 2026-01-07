package main

import "fmt"

type IndexNum int64

func main() {

	arr := []int{0, 1, 5, 7, 9, 11, 15, 20, 24, 35, 40, 45, 50, 55, 60, 65, 70, 75, 80, 85, 90, 95, 100}
	target := 24

	linearSearch(arr, target)
	binarySearch(arr, target)
	binarySearchRecursive(arr, target, 0, len(arr)-1)
}

func linearSearch(arr []int, target int) IndexNum {
	count := 0
	for i, v := range arr {
		count++
		if v == target {
			fmt.Println("Linear search count:", count)
			return IndexNum(i)
		}
	}

	return -1 // not found, return -1
}

func binarySearch(arr []int, target int) IndexNum {
	left, right := 0, len(arr)-1
	count := 0
	for left <= right {
		count++
		mid := (left + right) / 2
		if arr[mid] == target {
			fmt.Println("Binary search count:", count)
			return IndexNum(mid)
		} else if arr[mid] < target {
			left = mid + 1
		} else {
			right = mid - 1
		}
	}
	return -1
}

var binarySearchRecursiveCount = 0
func binarySearchRecursive(arr []int, target int, left int, right int) IndexNum {
	
	if left > right {
		return -1
	}
	binarySearchRecursiveCount++
	mid := (left + right) / 2
	if arr[mid] == target {
		fmt.Println("Binary search recursive count:", 	binarySearchRecursiveCount)
		return IndexNum(mid)
	} else if arr[mid] < target {
		return binarySearchRecursive(arr, target, mid+1, right)
	} else {
		return binarySearchRecursive(arr, target, left, mid-1)
	}
}
