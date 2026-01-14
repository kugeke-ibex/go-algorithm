package main

import (
	"fmt"
	"time"
)

func main() {
	memoizedFunc := memoize(longFunc)
	for i := range 10 {
		fmt.Println(memoizedFunc(i))
	}

	start := time.Now()
	for i := range 10 {
		fmt.Println(memoizedFunc(i))
	}
	elapsed := time.Since(start)
	fmt.Println(elapsed)
}

func longFunc(num int) int {
	r := 0
	for i := range 1000000 {
		r += num * i
	}

	return r
}

func memoize(fn func(int) int) func(int) int {
	cache := make(map[int]int)
	return func(num int) int {
		if val, ok := cache[num]; ok {
			return val
		}
		result := fn(num)
		cache[num] = result
		return result
	}
}
