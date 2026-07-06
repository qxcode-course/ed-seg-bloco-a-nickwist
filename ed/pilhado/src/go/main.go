package main

import (
	"bufio"
	"fmt"
	"os"
)

func podeAndar(mat [][]rune, linha, coluna int) bool {
	if linha < 0 || linha >= len(mat) {
		return false
	}

	if coluna < 0 || coluna >= len(mat[linha]) {
		return false
	}

	if mat[linha][coluna] == '#' || mat[linha][coluna] == '.' {
		return false
	}

	return mat[linha][coluna] == ' ' || mat[linha][coluna] == 'F'
}

func ehUmBeco(becos *Stack[Pos], pos Pos) bool {
	pilha := NewStack[Pos]()
	encontrou := false

	for !becos.IsEmpty() {
		p := becos.Pop()
		pilha.Push(p)

		if p.linha == pos.linha && p.coluna == pos.coluna {
			encontrou = true
		}
	}

	for !pilha.IsEmpty() {
		becos.Push(pilha.Pop())
	}

	return encontrou
}

func caminhar(mat [][]rune, inicio, fim Pos) {
	caminho := NewStack[Pos]()
	becos := NewStack[Pos]()

	caminho.Push(inicio)

	for !caminho.IsEmpty() {
		atual := caminho.Top()
		mat[atual.linha][atual.coluna] = '.'

		if atual.linha == fim.linha && atual.coluna == fim.coluna {
			break
		}

		validos := []Pos{}

		vizinhos := []Pos{
			{linha: atual.linha - 1, coluna: atual.coluna},
			{linha: atual.linha + 1, coluna: atual.coluna},
			{linha: atual.linha, coluna: atual.coluna + 1},
			{linha: atual.linha, coluna: atual.coluna - 1},
		}

		for _, vizinho := range vizinhos {
			if podeAndar(mat, vizinho.linha, vizinho.coluna) && !ehUmBeco(becos, vizinho) {
				validos = append(validos, vizinho)
			}
		}
		if len(validos) > 0 {
			caminho.Push(validos[0])
		} else {
			becos.Push(atual)
			caminho.Pop()

			if atual.linha != inicio.linha || atual.coluna != inicio.coluna {
				mat[atual.linha][atual.coluna] = ' '
			}
		}

	}
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	line := scanner.Text()
	var nl, nc int
	fmt.Sscanf(line, "%d %d", &nl, &nc)

	inicio := Pos{linha: -1, coluna: -1}
	fim := Pos{linha: -1, coluna: -1}

	grid := make([][]rune, 0, nl)

	for i := 0; i < nl; i++ {
		scanner.Scan()
		line := scanner.Text()

		linha := []rune(line)
		grid = append(grid, linha)

		for x, y := range linha {
			if y == 'I' {
				inicio = Pos{linha: i, coluna: x}
			}

			if y == 'F' {
				fim = Pos{linha: i, coluna: x}
			}
		}
	}
	caminhar(grid, inicio, fim)
	showGrid(grid)
}

func showGrid(grid [][]rune) {
	for _, line := range grid {
		fmt.Println(string(line))
	}
}
