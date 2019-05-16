package main

import "fmt"

const (
	players = 428
	turns   = 70825
)

func main() {
	scores := make(map[int]int)
	var board []int
	pos := 0
	for i := 0; i < turns+1; i++ {
		// fmt.Printf("%d: %v (%d / %d)\n", i, board, pos, len(board))
		if pos > len(board) {
			pos = 1
		}

		if i%23 == 0 && i > 0 {
			remPos := mod(pos-9, len(board))
			scores[i%players] += i
			scores[i%players] += board[remPos]
			board = append(board[:remPos], board[remPos+1:]...)
			pos = remPos + 2
			continue
		}

		if pos >= len(board) {
			board = append(board, i)
		} else {
			board = append(board[:pos], append([]int{i}, board[pos:]...)...)
		}

		pos = pos + 2
	}
	big := 0
	for _, score := range scores {
		if score > big {
			big = score
		}
	}
	fmt.Println(big)
}

func mod(d, m int) int {
	res := d % m
	if (res < 0 && m > 0) || (res > 0 && m < 0) {
		return res + m
	}
	return res
}
