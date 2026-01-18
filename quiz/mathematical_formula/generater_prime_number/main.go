package main

import (
	"fmt"
	"time"
)

func main() {
	start := time.Now()
	fmt.Printf("50 primes for generatePrimeV1: %v\n", generatePrimeV1(50))
	end := time.Now()
	fmt.Printf("Time taken: %v\n", end.Sub(start))

	start = time.Now()
	fmt.Printf("50 primes for generatePrimeV2: %v\n", generatePrimeV2(50))
	end = time.Now()
	fmt.Printf("Time taken: %v\n", end.Sub(start))
}

func generatePrimeV1(number int) []int {
	primes := []int{}
	var isPrime bool
	count := 0
	for x := 2;  x <= number; x++ {
		isPrime = true
		for y := 2; y < x; y++ {
			count++
			if x % y == 0 {
				isPrime = false
				break
			}
		}

		if isPrime {
			primes = append(primes, x)
		}
	}

	fmt.Printf("Count: %d\n", count)
	return primes
}

func generatePrimeV2(number int) []int {
	primes := []int{}
	cache := make(map[int]bool)

	count := 0
	for x := 2; x <= number; x++ {
		if isPrime, cached := cache[x]; cached && !isPrime {
			continue
		}

		primes = append(primes, x)
		cache[x] = true
		for y := x * 2; y <= number; y += x{
			count++
			cache[y] = false
		}
	}

	fmt.Printf("Count: %d\n", count)
	return primes
}