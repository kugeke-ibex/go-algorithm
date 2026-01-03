package main

import "fmt"

func main() {
	// O(1) < O(log(n)) < O(n) < O(n * log(n)) < O(n^2)
	fmt.Println("func1")
	fmt.Println(func1([]int{1, 2, 3, 4, 5}))
	fmt.Println("func2")
	func2(10)
	fmt.Println("func3")
	func3([]int{1, 2, 3, 4, 5})
	fmt.Println("func4")
	func4(10)
	fmt.Println("func5")
	func5([]int{1, 2, 3, 4, 5})
}

// O(1)
func func1(numbers []int) int {
	return numbers[0]
}

// O(log(n))
func func2(n int) {
	if n < 2 {
		return
	} else {
		fmt.Println(n)
		func2(n/2)
	}
}

// O(n)
func func3(numbers []int) {
	for _, v := range numbers {
		fmt.Println(v)
	}
}

// O(n * log(n))
func func4(n int) {
	for i := range n {
		fmt.Print(i, "　")
	}
	if n < 2 {
		return
	}
	func4(n/2)
}

// O(n^2)
func func5(numbers []int) {
	for i := 0; i < len(numbers); i++ {
		for j := 0; j < len(numbers); j++ {
			fmt.Println(numbers[i], numbers[j])
		}
	}
}