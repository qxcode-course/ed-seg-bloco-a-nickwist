
package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func tostr(vet []int) string {

	if len(vet) == 0 {
		return ""
	}

	if len(vet) == 1 {
		return fmt.Sprintf("%d",vet[0])
	}

	return fmt.Sprintf("%d, ", vet[0] ) + tostr(vet[1:])
	
}

func tostrrev(vet []int) string {
	if len(vet) == 0 {
		return ""
	}

	if len(vet) == 1 {
		return fmt.Sprintf("%d", vet[len(vet)-1])
	}

	return fmt.Sprintf("%d, ", vet[(len(vet))-1] ) + tostrrev(vet[:len(vet)-1])
	
}

// reverse: inverte os elementos do slice
func reverse(vet []int) {
	if len(vet) <= 1 {
		return
	}
	
	vet[0], vet[len(vet)-1] = vet[len(vet)-1], vet[0]
	
	reverse(vet[1 : len(vet)-1])

}

// sum: soma dos elementos do slice
func sum(vet []int) int {
	if len(vet) == 0 {
		return 0
	}

	if len(vet) == 1 {
		return vet[0]
	}

	return vet[0] + sum(vet[1:])
}

// mult: produto dos elementos do slice
func mult(vet []int) int {
	if len(vet) == 0 {
		return 1
	}

	if len(vet) == 1 {
		return vet[0]
	}

	return vet[0] * mult(vet[1:])
}

// min: retorna o índice e valor do menor valor
// crie uma função recursiva interna do modelo
// var rec func(v []int) (int, int)
// para fazer uma recursão que retorna valor e índice
func min(vet []int) int {
	
	var rec func(v []int) (int, int) 

	rec = func(v []int) (int, int) {
		
		if len(v) == 0 {
			return -1, 0
		}
		
		if len(v) == 1 {
			return 0, v[0]
		}

		index, menor := rec(v[1:])

		if v[0] <= menor {
           	return 0, v[0]
        }

    	return index + 1, menor
	}

	indice, _ := rec(vet)
	
	return indice
	}


func main() {
	var vet []int
	scanner := bufio.NewScanner(os.Stdin)
	for {
		if !scanner.Scan() {
			break
		}
		line := scanner.Text()
		args := strings.Fields(line)
		fmt.Println("$" + line)

		switch args[0] {
		case "end":
			return
		case "read":
			vet = nil
			for _, arg := range args[1:] {
				if val, err := strconv.Atoi(arg); err == nil {
					vet = append(vet, val)
				}
			}
		case "tostr":
			fmt.Println("[" + tostr(vet) + "]")
		case "torev":
			fmt.Println("[" + tostrrev(vet) + "]")
		case "reverse":
			reverse(vet)
		case "sum":
			fmt.Println(sum(vet))
		case "mult":
			fmt.Println(mult(vet))
		case "min":
			fmt.Println(min(vet))
		default:
			fmt.Println("fail: comando invalido")
		}
	}
}
