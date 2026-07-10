package main

import (
	"bufio"
	"fmt"
	"os"
)

func DFS(board [][]byte, i, j int) {
	if i < 0 || j < 0 || i >= len(board) || j >= len(board[0]) {
		return
	}

	if board[i][j] != 'O' {
		return
	}

	board[i][j] = '.'

	DFS(board, i+1, j)
	DFS(board, i-1, j)
	DFS(board, i, j+1)
	DFS(board, i, j-1)

}

// NÃO ALTERE A ASSINATURA DA FUNÇÃO solve
func solve(board [][]byte) {
	for i := 0; i < len(board); i++ {
		DFS(board, i, 0)
		DFS(board, i, len(board[0])-1)
	}

	for j := 0; j < len(board[0]); j++ {
		DFS(board, 0, j)
		DFS(board, len(board)-1, j)
	}

	for i := 0; i < len(board); i++ {
		for j := 0; j < len(board[0]); j++ {
			if board[i][j] == 'O' {
				board[i][j] = 'X'
			} else if board[i][j] == '.' {
				board[i][j] = 'O'
			}
		}
	}
}

// NÃO ALTERE A MAIN
func main() {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	var nrows, ncols int
	fmt.Sscanf(scanner.Text(), "%d %d", &nrows, &ncols)
	board := make([][]byte, nrows)
	for i := 0; i < nrows; i++ {
		scanner.Scan()
		board[i] = []byte(scanner.Text())
	}
	solve(board)
	for _, row := range board {
		fmt.Println(string(row))
	}
}
