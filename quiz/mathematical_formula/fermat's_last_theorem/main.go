package main

import (
	"fmt"
	"math"
	"time"
)

func main() {
	fmt.Printf("Fermat's Last Theorem V1(10, 2): %v\n", fermatLastTheoremV1(10, 2))
	fmt.Printf("Fermat's Last Theorem V1(20, 2): %v\n", fermatLastTheoremV1(20, 2))
	fmt.Printf("Fermat's Last Theorem V1(10, 3): %v\n", fermatLastTheoremV1(10, 3))
	fmt.Printf("Fermat's Last Theorem V2(10, 2): %v\n", fermatLastTheoremV2(10, 2))
	fmt.Printf("Fermat's Last Theorem V2(20, 2): %v\n", fermatLastTheoremV2(20, 2))
	fmt.Printf("Fermat's Last Theorem V2(10, 3): %v\n", fermatLastTheoremV2(10, 3))

	fmt.Println("--------------------------------")
	startTime := time.Now()
	fmt.Printf("Fermat's Last Theorem V1(10000, 2): %v\n", fermatLastTheoremV1(100, 2))
	endTime := time.Now()
	fmt.Printf("Time taken: %v\n", endTime.Sub(startTime))

	startTime = time.Now()
	fmt.Printf("Fermat's Last Theorem V2(10000, 2): %v\n", fermatLastTheoremV2(100, 2))
	endTime = time.Now()
	fmt.Printf("Time taken: %v\n", endTime.Sub(startTime))

	fmt.Println("--------------------------------")
	for i := 2; i <= 100; i++ {
		fmt.Printf("Fermat's Last Theorem V1(20, %d): %v\n", i, fermatLastTheoremV1(20, i))
		fmt.Printf("Fermat's Last Theorem V2(20, %d): %v\n", i, fermatLastTheoremV2(20, i))
	}
}

func fermatLastTheoremV1(maxNumber int, squireNumber int) [][]int {
	result := make([][]int, 0)
	if squireNumber < 2 {
		return result
	}

	maxZ := int(math.Pow(math.Pow(float64(maxNumber-1), 2.0)+math.Pow(float64(maxNumber), 2.0), 1.0/float64(squireNumber)))
	for x := 1; x <= maxNumber; x++ {
		for y := x + 1; y <= maxNumber; y++ {
			for z := y + 1; z <= maxZ; z++ {
				if math.Pow(float64(x), float64(squireNumber))+math.Pow(float64(y), float64(squireNumber)) == math.Pow(float64(z), float64(squireNumber)) {
					result = append(result, []int{x, y, z})
				}
			}
		}
	}
	return result
}

func fermatLastTheoremV2(maxNumber int, squireNumber int) [][]int {
	result := make([][]int, 0)
	if squireNumber < 2 {
		return result
	}

	for x := 1; x <= maxNumber; x++ {
		for y := x + 1; y <= maxNumber; y++ {
			powSum := math.Pow(float64(x), float64(squireNumber)) + math.Pow(float64(y), float64(squireNumber))
			if powSum > math.MaxInt64 {
				continue
			}

			z := math.Pow(powSum, 1.0/float64(squireNumber))
			if z != float64(int(z)) {
				continue
			}

			zInt := int(z)

			if zInt <= y {
				continue
			}

			zPow := math.Pow(float64(zInt), float64(squireNumber))
			if zPow == powSum {
				result = append(result, []int{x, y, zInt})
			}
		}
	}
	return result
}
