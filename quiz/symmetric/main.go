package main

import (
	"fmt"
)

func main() {
	input := [][2]int{
		{1, 2},
		{3, 5},
		{4, 7},
		{5, 3},
		{7, 4},
	}

	findPairs := findPair(input)
	fmt.Println(findPairs)
}

func findPair(list [][2]int) [][2]int {
	findPairs := make([][2]int, 0)

	cache := make(map[int]int)
	for _, pair := range list {
		first, second := pair[0], pair[1]

		if _, ok := cache[second]; !ok {
			cache[first] = second
		} else if first == cache[second]{
			findPairs = append(findPairs, [2]int{first, second})
		}
	}

	return findPairs
}