package main

import (
	"bufio"
	"fmt"
	"os"
)

type Pos struct {
	l, c int
}

func (p Pos) getNeig() []Pos {
	return nil
}

func inside(grid [][]rune, pos Pos) bool {
	nrows := len(grid)
	ncols := len(grid[0])
	return pos.l >= 0 && pos.l < nrows && pos.c >= 0 && pos.c < ncols
}

func match(grid [][]rune, pos Pos, char rune) bool {
	return inside(grid, pos) && grid[pos.l][pos.c] == char
}

func search(grid [][]rune, startPos Pos, endPos Pos) {
	fila := []Pos{startPos}

	visitado := make(map[Pos]bool)
	visitado[startPos] = true

	anterior := make(map[Pos]Pos)

	direcoes := []Pos{
		{1, 0},
		{-1, 0},
		{0, 1},
		{0, -1},
	}

	for len(fila) > 0 {
		atual := fila[0]
		fila = fila[1:]

		if atual == endPos {
			voltar(grid, anterior, startPos, endPos)
			break
		}

		for _, direcao := range direcoes {
			vizinho := Pos{
				l: atual.l + direcao.c,
				c: atual.c + direcao.l,
			}

			if inside(grid, vizinho) && grid[vizinho.l][vizinho.c] != '#' && !visitado[vizinho] {
				visitado[vizinho] = true
				anterior[vizinho] = atual
				fila = append(fila, vizinho)
			}

		}
	}
}

func voltar(grid [][]rune, anterior map[Pos]Pos, startPos, endPos Pos) {
	atual := endPos

	for atual != startPos {
		grid[atual.l][atual.c] = '.'
		atual = anterior[atual]
	}

	grid[startPos.l][startPos.c] = '.'
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)

	var nl, nc int
	scanner.Scan()
	line := scanner.Text()
	fmt.Sscanf(line, "%d %d", &nl, &nc)
	mat := make([][]rune, nl) // Inicializa a matriz de runes

	// Carregando matriz
	for i := range nl {
		scanner.Scan()
		line := scanner.Text()
		mat[i] = []rune(line)
	}

	var inicio, fim Pos

	// Procurando inicio e fim e colocando ' ' nas posições iniciais
	for l := range nl {
		for c := range nc {
			if mat[l][c] == 'I' {
				mat[l][c] = ' '
				inicio = Pos{l, c}
			}
			if mat[l][c] == 'F' {
				mat[l][c] = ' '
				fim = Pos{l, c}
			}
		}
	}

	search(mat, inicio, fim)

	for _, line := range mat {
		fmt.Println(string(line)) // Converte o slice de runes de volta para string para imprimir
	}
}
