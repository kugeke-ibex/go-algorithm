package main

import (
	"fmt"
	"strconv"
	"strings"
)

func main() {

	triangle := generatePascalTriangle(12)
	printPascalTriangle(triangle)
}

func generatePascalTriangle(depth int) [][]int {
	triangle := make([][]int, depth)
	for i := 0; i < depth; i++ {
		triangle[i] = make([]int, i+1)
		triangle[i][0] = 1
		triangle[i][i] = 1
		for j := 1; j < i; j++ {
			triangle[i][j] = triangle[i-1][j-1] + triangle[i-1][j]
		}
	}
	return triangle
}

func printPascalTriangle(triangle [][]int) {
	if len(triangle) == 0 {
		return
	}

	// 最後の行の最大数字を見つけて、最大桁数を計算
	maxValue := 0
	for _, value := range triangle[len(triangle)-1] {
		if value > maxValue {
			maxValue = value
		}
	}
	maxDigitWidth := len(strconv.Itoa(maxValue))

	// 最後の行の幅（各数字の幅 + スペース）を計算
	lastRowWidth := len(triangle[len(triangle)-1])*maxDigitWidth + (len(triangle[len(triangle)-1]) - 1)

	for _, row := range triangle {
		// 各行の数字を文字列に変換し、パディング
		rowStr := make([]string, len(row))
		for i, value := range row {
			rowStr[i] = fmt.Sprintf("%*d", maxDigitWidth, value)
		}

		// 行を結合
		rowLine := strings.Join(rowStr, " ")

		// 中心に揃えるためのインデントを計算
		currentWidth := len(rowLine)
		indent := (lastRowWidth - currentWidth) / 2

		// インデントを追加して出力
		fmt.Print(strings.Repeat(" ", indent))
		fmt.Println(rowLine)
	}
}
