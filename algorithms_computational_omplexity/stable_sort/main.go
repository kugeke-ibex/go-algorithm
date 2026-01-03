package main

import "fmt"

func main() {

	// 安定ソートとは: ソート判定において同一であると判断された入力データの順序がソート後も変わらないソートのこと
	stableList := []map[int]string{
		{
			1: "John",
		},
		{
			5: "Jane",
		},
		{
			2: "Jim",
		},
		{
			4: "Jill",
		},
		{
			2: "Jimmy",
		},
	}

	unstableList := []map[int]string{
		{
			1: "John",
		},
		{
			5: "Jane",
		},
		{
			2: "Jim",
		},
		{
			4: "Jill",
		},
		{
			2: "Jimmy",
		},
	}

	fmt.Println("stable")
	fmt.Println(stable(stableList))
	fmt.Println("unstable")
	fmt.Println(unstable(unstableList))
	// stable
	// [map[1:John] map[2:Jim] map[2:Jimmy] map[4:Jill] map[5:Jane]]
	// unstable → JimとJimmyの順序が変わってしまっている
	// [map[1:John] map[2:Jimmy] map[2:Jim] map[4:Jill] map[5:Jane]]
}

func stable(numbers []map[int]string) []map[int]string {
	numbers[1], numbers[2] = numbers[2], numbers[1]
	numbers[2], numbers[4] = numbers[4], numbers[2]
	return numbers
}

func unstable(numbers []map[int]string) []map[int]string {
	numbers[1], numbers[4] = numbers[4], numbers[1]
	return numbers
}
