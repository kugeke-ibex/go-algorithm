package main

import (
	"fmt"
)

func main() {
	for _, perm := range allPermutations([]int{1, 2, 3}) {
		fmt.Println(perm)
	}
}

func allPermutations(elements []int) [][]int {
	result := make([][]int, 0)

	if len(elements) <= 1 {
		return [][]int{elements}
	}

	fmt.Printf("elements[1:]: %v\n", elements[1:])
	for _, perm := range allPermutations(elements[1:]) {
		fmt.Printf("perm: %v\n", perm)
		for i := range len(elements) {
			// copy関数を使ってrest[:i]のコピーを作成
			copied := make([]int, i)
			copy(copied, perm[:i])
			temp := append(copied, elements[0:1]...)
			temp = append(temp, perm[i:]...)
			result = append(result, temp)
			fmt.Printf("result: %v\n", result)
		}
	}

	return result
}
