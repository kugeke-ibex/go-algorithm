package main

import (
	"fmt"
	"strings"
)

func main() {
	w := []string{"h", "y", "n", "p", "t", "o"}
	indexes := []int{3, 1, 5, 0, 2, 4}
	fmt.Printf("ordered v1 by indexes: %s\n", orderChangeByIndexesV1(w, indexes))
	fmt.Printf("ordered v2 by indexes: %s\n", orderChangeByIndexesV2(w, indexes))
}


func orderChangeByIndexesV1(chars []string, indexes []int) string {
	temp := make([]string, len(chars))
	for i, index := range indexes {
		temp[index] = chars[i]
	}

	return strings.Join(temp, "")
}

func orderChangeByIndexesV2(chars []string, indexes []int) string {
	i, length := 0, len(chars) - 1
	for i < length {
		for i != indexes[i] {
			index := indexes[i]
			chars[i], chars[index] = chars[index], chars[i]
			indexes[i], indexes[index] = indexes[index], indexes[i]
		}
		i++
	}

	return strings.Join(chars, "")
}

