package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

type Pair struct {
	One int
	Two int
}

func occurr(vet []int) []Pair {
	repetidos := make(map[int]int)

	for _, v := range vet {
		if v < 0 {
			v = -v
		}

		repetidos[v]++
	}

	pares := make([]Pair, 0, len(repetidos))

	for v, qntd := range repetidos {
		pares = append(pares, Pair{
			One: v,
			Two: qntd,
		})
	}

	sort.Slice(pares, func(v, qntd int) bool {
		return pares[v].One < pares[qntd].One
	})

	return pares

}

func teams(vet []int) []Pair {
	times := []Pair{}
	if len(vet) == 0 {
		return nil
	}

	atual := vet[0]
	count := 1

	for i := 1; i < len(vet); i++ {
		if vet[i] == atual {
			count++
		} else {
			times = append(times, Pair{One: atual, Two: count})
			atual = vet[i]
			count = 1
		}

	}
	times = append(times, Pair{One: atual, Two: count})

	return times
}

func mnext(vet []int) []int {
	mulherprox := []int{}

	if len(vet) == 1 {
			return []int{0}
		} 

	for i := 0; i < len(vet); i++ {
		if vet[i] > 0 {
			if i == 0 {
				if vet[i+1] < 0 {
					mulherprox = append(mulherprox, 1)
				} else {
					mulherprox = append(mulherprox, 0)
				} 
			} else if i == (len(vet)-1) {
				if vet[i-1] < 0 {
					mulherprox = append(mulherprox, 1)
				} else {
					mulherprox = append(mulherprox, 0)
				}
			} else {
				if vet[i+1] < 0 || vet[i-1] < 0 {
					mulherprox = append(mulherprox, 1)
				} else {
					mulherprox = append(mulherprox, 0)
				}
			}
		} else {
			mulherprox = append(mulherprox, 0)
		}
	}
		return mulherprox
	
}

func alone(vet []int) []int {
		semMulherProx := []int{}

	if len(vet) == 1 && vet[0] > 0 {
			return []int{1}
	} 

	for i := 0; i < len(vet); i++ {
		if vet[i] > 0 {
			if i == 0 {
				if vet[i+1] < 0 {
					semMulherProx = append(semMulherProx, 0)
				} else {
					semMulherProx = append(semMulherProx, 1)
				} 
			} else if i == (len(vet)-1) {
				if vet[i-1] < 0 {
					semMulherProx = append(semMulherProx, 0)
				} else {
					semMulherProx = append(semMulherProx, 1)
				}
			} else {
				if vet[i+1] < 0 || vet[i-1] < 0 {
					semMulherProx = append(semMulherProx, 0)
				} else {
					semMulherProx = append(semMulherProx, 1)
				}
			}
		} else {
			semMulherProx = append(semMulherProx, 0)
		}
	}
		return semMulherProx
	
}

func couple(vet []int) int {
	casais := 0;

	if len(vet) == 1 {
		return 0
	}

	for i := 0; i < len(vet); i++ {
		for j := i + 1; j < len(vet); j++ {
			if vet[i] == -(vet[j]) || vet[j] == -(vet[i]) {
				casais++
				i+=2
			}
		}
	}

	return casais
}

	func hasSubseq(vet []int, seq []int, pos int) bool {
		
		if pos+len(seq) > len(vet) {
        return false
    }

    for j := 0; j < len(seq); j++ {
        if vet[pos+j] != seq[j] {
            return false
        }
    }

    return true

	}

	func subseq(vet []int, seq []int) int {
		pos := -1

		for i := 0; i < len(vet); i++ {
			if hasSubseq(vet, seq, i) {
				pos = i
				break
			}
		}

		return pos
	}

func erase(vet []int, posList []int) []int {
	listaNova := []int{}
	
	for i := 0; i < len(vet); i++ {
		achou := false

		for j:= 0; j < len(posList); j++ {
			if i == posList[j]{
				achou = true
				break
			} 
		}

		if !achou {
			listaNova = append(listaNova, vet[i])
		} 
	} 

	return listaNova
}

func clear(vet []int, value int) []int {
	listaNova := []int{}

	for _, v := range vet {
		if v != value {
			listaNova = append(listaNova, v)
		}
	}
	return listaNova
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	for {
		fmt.Print("$")
		if !scanner.Scan() {
			break
		}
		line := scanner.Text()
		args := strings.Split(line, " ")
		fmt.Println(line)

		switch args[0] {
		case "occurr":
			printSlice(occurr(str2vet(args[1])))
		case "teams":
			printSlice(teams(str2vet(args[1])))
		case "mnext":
			printSlice(mnext(str2vet(args[1])))
		case "alone":
			printSlice(alone(str2vet(args[1])))
		case "erase":
			printSlice(erase(str2vet(args[1]), str2vet(args[2])))
		case "clear":
			val, _ := strconv.Atoi(args[2])
			printSlice(clear(str2vet(args[1]), val))
		case "subseq":
			fmt.Println(subseq(str2vet(args[1]), str2vet(args[2])))
		case "couple":
			fmt.Println(couple(str2vet(args[1])))
		case "end":
			return
		default:
			fmt.Println("Invalid command")
		}
	}
}

// Funções auxiliares

func str2vet(str string) []int {
	if str == "[]" {
		return nil
	}
	str = str[1 : len(str)-1]
	parts := strings.Split(str, ",")
	var vet []int
	for _, part := range parts {
		num, _ := strconv.Atoi(strings.TrimSpace(part))
		vet = append(vet, num)
	}
	return vet
}

func printSlice[T any](vet []T) {
	fmt.Print("[")
	for i, x := range vet {
		if i > 0 {
			fmt.Print(", ")
		}
		fmt.Print(x)
	}
	fmt.Println("]")
}

func (p Pair) String() string {
	return fmt.Sprintf("(%v, %v)", p.One, p.Two)
}
