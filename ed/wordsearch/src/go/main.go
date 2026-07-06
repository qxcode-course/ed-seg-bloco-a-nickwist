package main

import (
	"bufio"
	"fmt"
	"os"
)

func buscar(grid [][]byte, word string, i, j, posatual int) bool {
	if posatual == len(word) {
		return true
	}

	if i < 0 || i >= len(grid) {
		return false
	}

	if j < 0 || j >= len(grid[i]) {
		return false
	}

	if grid[i][j] != word[posatual] {
		return false
	}

	voltar := grid[i][j]
	grid[i][j] = '#'

	if buscar(grid, word, i+1, j, posatual+1) ||
		buscar(grid, word, i-1, j, posatual+1) ||
		buscar(grid, word, i, j+1, posatual+1) ||
		buscar(grid, word, i, j-1, posatual+1) {
		return true
	}
	grid[i][j] = voltar
	return false
}

// Não mude a assinatura desta função, ela é a função chamada pelo LeetCode
func exist(grid [][]byte, word string) bool {
	for i := 0; i < len(grid); i++ {
		for j := 0; j < len(grid[i]); j++ {
			if grid[i][j] == word[0] {
				if buscar(grid, word, i, j, 0) {
					return true
				}
			}
		}
	}

	return false

}

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	var word string
	fmt.Sscanf(scanner.Text(), "%s", &word)
	grid := make([][]byte, 0)
	for scanner.Scan() {
		grid = append(grid, []byte(scanner.Text()))
	}
	if exist(grid, word) {
		fmt.Println("true")
	} else {
		fmt.Println("false")
	}
}
