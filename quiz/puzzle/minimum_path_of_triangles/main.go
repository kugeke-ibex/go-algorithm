package main

import (
	"fmt"
	"math/rand"
	"slices"
	"strconv"
	"strings"
)

func main() {
	triangle := generateTriangleList(5, 20)
	printTriangle(triangle)
	minPath := sumMinPaht(triangle)
	fmt.Println("Minimum path sum:", *minPath)
}

func sumMinPaht(triangle [][]int) *int {
	treeSum := make([][]int, len(triangle))
	copy(treeSum, triangle)
	j, lengthTrianggle := 1, len(triangle)
	if lengthTrianggle < 1 {
		return nil
	}
	for j < lengthTrianggle {
		line := triangle[j]
		linePathSum := make([]int, 0, len(line))
		for i, value := range line {
			var sumValue int
			if i == 0 {
				sumValue = line[i] + treeSum[j-1][0]
			} else if i == len(line)-1 {
				sumValue = line[i] + treeSum[j-1][i-1]
			} else {
				minValue := min(treeSum[j-1][i-1], treeSum[j-1][i])
				sumValue = value + minValue
			}
			linePathSum = append(linePathSum, sumValue)
		}
		treeSum[j] = linePathSum
		j++
	}

	printTriangle(treeSum)

	minValue := slices.Min(treeSum[lengthTrianggle-1])
	return &minValue
}

func generateTriangleList(depth int, maxNumber int) [][]int {
	triangle := make([][]int, depth)
	for i := 0; i < depth; i++ {
		triangle[i] = make([]int, i+1)
		for j := 0; j < i+1; j++ {
			triangle[i][j] = rand.Intn(maxNumber)
		}
	}
	return triangle
}

func printTriangle(triangle [][]int) {
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
