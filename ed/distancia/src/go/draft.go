package main

import (
	"fmt"
)

func podeColocar(seq []rune, pos, digito, limite int) bool {
	char := rune('0' + digito)

	for i := 0; i <= limite; i++ {
		esquerda := pos - i
		direita := pos + i

		if esquerda >= 0 && seq[esquerda] == char {
			return false
		}

		if direita < len(seq) && seq[direita] == char {
			return false
		}
	}

	return true
}

func substituicao(L, inicio int, seq []rune) bool {

	for i := inicio; i < len(seq); i++ {
		if seq[i] == '.' {
			for digito := 0; digito <= L; digito++ {
				if podeColocar(seq, i, digito, L) {
					seq[i] = rune('0' + digito)
				
				    if substituicao(L, i+1, seq) {
					    return true
				    }
                }
                seq[i] = '.'
            }
            return false
		}
	}
    fmt.Println(string(seq))
    return true
}

func main() {
	var seq string
	var L int

	fmt.Scan(&seq)
	fmt.Scan(&L)

	seqRune := []rune(seq)

	substituicao(L, 0, seqRune)
}
