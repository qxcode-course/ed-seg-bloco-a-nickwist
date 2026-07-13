package main

import "fmt"

func apodrecendo(matriz [][]int) int {
	if len(matriz) == 0 || len(matriz[0]) == 0 {
		return 0
	}

	filas := len(matriz)
	colunas := len(matriz[0])

	fila := [][2]int{}
	laranjasfrescas := 0

	for i := 0; i < filas; i++ {
		for j := 0; j < colunas; j++ {
			if matriz[i][j] == 2 {
				fila = append(fila, [2]int{i, j})
			} else if matriz[i][j] == 1 {
				laranjasfrescas++
			}
		}
	}

	direcoes := [][2]int{
		{1, 0},
		{-1, 0},
		{0, 1},
		{0, -1},
	}

	minutos := 0

	for len(fila) > 0 && laranjasfrescas > 0 {
		tamanho := len(fila)

		for i := 0; i < tamanho; i++ {
			atual := fila[0]
			fila = fila[1:]

			fileira := atual[0]
			coluna := atual[1]

			for _, direcao := range direcoes {
				novafila := fileira + direcao[0]
				novacoluna := coluna + direcao[1]

				if novafila >= 0 && novacoluna >= 0 && novacoluna < colunas && novafila < filas && matriz[novafila][novacoluna] == 1 {
					matriz[novafila][novacoluna] = 2
					laranjasfrescas--

					fila = append(fila, [2]int{novafila, novacoluna})
				}
			}
		}

		minutos++

	}

	if laranjasfrescas > 0 {
		return -1
	}

	return minutos

}

func main() {
	var colunas, linhas int

	fmt.Scan(&colunas, &linhas)

	matriz := make([][]int, linhas)

	for i := 0; i < linhas; i++ {
		matriz[i] = make([]int, colunas)

		for j := 0; j < colunas; j++ {
			fmt.Scan(&matriz[i][j])
		}
	}

	fmt.Println(apodrecendo(matriz))

}
