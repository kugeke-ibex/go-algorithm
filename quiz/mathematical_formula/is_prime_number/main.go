package main

import (
	"fmt"
	"math"
	"math/rand"
	"time"
)

func main() {
	fmt.Printf("Is 10 a prime number (V1)? %v\n", isPrimeV1(10))
	fmt.Printf("Is 11 a prime number (V1)? %v\n", isPrimeV1(11))
	fmt.Printf("Is 10 a prime number (V2)? %v\n", isPrimeV2(10))
	fmt.Printf("Is 11 a prime number (V2)? %v\n", isPrimeV2(11))
	fmt.Printf("Is 10 a prime number (V3)? %v\n", isPrimeV3(10))
	fmt.Printf("Is 11 a prime number (V3)? %v\n", isPrimeV3(11))
	fmt.Printf("Is 10 a prime number (V4)? %v\n", isPrimeV4(10))
	fmt.Printf("Is 11 a prime number (V4)? %v\n", isPrimeV4(11))
	fmt.Println("--------------------------------")

	number := rand.Intn(100)
	startTime := time.Now()
	fmt.Printf("Is %d a prime number (V1)? %v\n", number, isPrimeV1(number))
	elapsedTime := time.Since(startTime)
	fmt.Printf("Time taken: %v\n", elapsedTime)

	startTime = time.Now()
	fmt.Printf("Is %d a prime number (V2)? %v\n", number, isPrimeV2(number))
	elapsedTime = time.Since(startTime)
	fmt.Printf("Time taken: %v\n", elapsedTime)

	startTime = time.Now()
	fmt.Printf("Is %d a prime number (V3)? %v\n", number, isPrimeV3(number))
	elapsedTime = time.Since(startTime)
	fmt.Printf("Time taken: %v\n", elapsedTime)

	startTime = time.Now()
	fmt.Printf("Is %d a prime number (V4)? %v\n", number, isPrimeV4(number))
	elapsedTime = time.Since(startTime)
	fmt.Printf("Time taken: %v\n", elapsedTime)
}

func isPrimeV1(number int) bool {
	if number <= 1 {
		return false
	}

	for x := 2; x < number; x++ {
		if number % x == 0 {
			return false
		}
	}

	return true
}

func isPrimeV2(number int) bool {
	/**
	36 = 1 * 36
	36 = 2 * 18
	36 = 3 * 12
	36 = 4 * 9
	36 = 6 * 6 <= √36 = 6 → これより大きい数は判定不要
	36 = 9 * 4
	36 = 12 * 3
	36 = 18 * 2
	36 = 36 * 1
	*/

	if number <= 1 {
		return false
	}

	for x := 2; x < int(math.Sqrt(float64(number))) + 1; x++ {
		if number % x == 0 {
			return false
		}
	}

	return true
}

func isPrimeV3(number int) bool {
	if number <= 1 {
		return false
	}
	if number == 2 {
		return true
	}
	if number % 2 == 0 {
		return false
	}

	// math.Sqrt不要の実装
	i := 2
	for  i * i <= number {
		if number % i == 0 {
			return false
		}
		i++
	}

	return true
}

func isPrimeV4(number int) bool {
	if number <= 1 {
		return false
	}

	if number <= 3 {
		return true
	}

	if number % 2 == 0 || number % 3 == 0 {
		return false
	}

	i := 5
	for i * i <= number {
		if number % i == 0 || number % (i + 2) == 0 {
			return false
		}
		i += 6
	}

	return true
}