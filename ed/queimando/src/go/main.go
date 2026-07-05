package main

import (
	"bufio"
	"fmt"
	"os"
)

func podequeimar(mat [][]rune, l, c int) bool {
	if l < 0 || l >= len(mat) {
		return false
	}

	if c < 0 || c >= len(mat[l]) {
		return false
	}

	return mat[l][c] == '#'

}

func burnTrees(grid [][]rune, l, c int) {
	stack := NewStack[Pos]()

	stack.Push(Pos{linha: l, coluna: c})

	for !stack.IsEmpty() {
		pos := stack.Pop()

		if !podequeimar(grid, pos.linha, pos.coluna) {
			continue
		}

		grid[pos.linha][pos.coluna] = 'o'

		stack.Push(Pos{linha: pos.linha + 1, coluna: pos.coluna})
		stack.Push(Pos{linha: pos.linha - 1, coluna: pos.coluna})
		stack.Push(Pos{linha: pos.linha, coluna: pos.coluna - 1})
		stack.Push(Pos{linha: pos.linha, coluna: pos.coluna + 1})
	}

}

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	line := scanner.Text()
	var nl, nc, lfire, cfire int
	fmt.Sscanf(line, "%d %d %d %d", &nl, &nc, &lfire, &cfire)

	grid := make([][]rune, 0, nl)
	for range nl {
		scanner.Scan()
		line := []rune(scanner.Text())
		grid = append(grid, line)
	}
	burnTrees(grid, lfire, cfire)
	showGrid(grid)
}

func showGrid(mat [][]rune) {
	for _, linha := range mat {
		fmt.Println(string(linha))
	}
}
