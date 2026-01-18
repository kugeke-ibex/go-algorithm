package main

import (
	"fmt"
)

func main() {
	hanoi(3, "A", "C", "B")

	fmt.Println("--------------------------------")
	moves := getHanoiMoves(4, "A", "C", "B")
	for _, move := range moves {
		fmt.Println(move)
	}
}

func hanoi(disk int, src string, dest string, support string) {
	if disk < 1 {
		return
	}

	hanoi(disk-1, src, support, dest)
	fmt.Printf("Move %d from %s to %s\n", disk, src, dest)
	hanoi(disk-1, support, dest, src)
}

func _hanoi(move *[][]string, disk int, src string, dest string, support string) {
	if disk < 1 {
		return
	}

	_hanoi(move, disk-1, src, support, dest)
	*move = append(*move, []string{fmt.Sprintf("%d", disk), src, dest})
	_hanoi(move, disk-1, support, dest, src)
}

func getHanoiMoves(disk int, src string, dest string, support string) [][]string {
	moves := make([][]string, 0)

	_hanoi(&moves, disk, src, dest, support)
	return moves
}
