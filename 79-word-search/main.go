package main

import "fmt"

func main() {
	var (
		board     [][]byte
		boardOrig [][]string
		word      string
		exists    bool
	)

	board = [][]byte{{'A', 'B', 'C', 'E'}, {'S', 'F', 'C', 'S'}, {'A', 'D', 'E', 'E'}}
	boardOrig = [][]string{{"A", "B", "C", "E"}, {"S", "F", "C", "S"}, {"A", "D", "E", "E"}}

	word = "ABCCED"
	exists = exist(board, word)
	fmt.Printf("The word <%s> exists in board %v? %t\n", word, boardOrig, exists)

	word = "SEE"
	exists = exist(board, word)
	fmt.Printf("The word <%s> exists in board %v? %t\n", word, boardOrig, exists)

	word = "ABCB"
	exists = exist(board, word)
	fmt.Printf("The word <%s> exists in board %v? %t\n", word, boardOrig, exists)

	word = "SFCC"
	exists = exist(board, word)
	fmt.Printf("The word <%s> exists in board %v? %t\n", word, boardOrig, exists)
}

func exist(board [][]byte, word string) bool {
	rows, cols := len(board), len(board[0])

	var dfs func(row, col, position int) bool
	dfs = func(row, col, position int) bool {
		if row < 0 || col < 0 ||
			row >= rows || col >= cols ||
			position == len(word) {
			return false
		}

		if word[position] != board[row][col] || board[row][col] == '*' {
			return false
		}

		if position == len(word) - 1 {
			return true
		}

		temp := board[row][col]
		board[row][col] = '*'
		found := (dfs(row + 1, col, position + 1) ||
			dfs(row - 1, col, position + 1) ||
			dfs(row, col + 1, position + 1) ||
			dfs(row, col - 1, position + 1))
		board[row][col] = temp

		return found
	}

	for r := 0; r < rows; r++ {
		for c := 0; c < cols; c++ {
			if dfs(r, c, 0) {
				return true
			}
		}
	}
	return false
}
