package main

import "fmt"

func main() {
	//fmt.Println(hasWin([]int{1, 1, 1}))
	fmt.Println(leftMost([]int{0, 1, 1}))
	fmt.Println(rightMost([]int{0, 1, 0}))
	fmt.Println(rightMost([]int{1, 1, 0}))
	fmt.Println(middle([]int{0, 0, 0}))
}

func hasWin(board []int) bool {
	score := 0
	for _, v := range board {
		score += v
	}
	return score == 3
}

func leftMost(board []int) bool {
	board[0] = 1
	return hasWin(board)
}

func rightMost(board []int) bool {
	board[2] = 1
	return hasWin(board)
}

func middle(board []int) bool {
	board[1] = 1
	return hasWin(board)
}
