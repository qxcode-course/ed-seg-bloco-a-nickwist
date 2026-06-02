package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func getMen(vet []int) []int {
	homens := []int{}
	for i := 0; i < len(vet); i++ {
		if vet[i] > 0 {
			homens = append(homens, vet[i])
		}
	}
	return homens
}

func getCalmWomen(vet []int) []int {
	mulhercalma := []int{}
	for i := 0; i < len(vet); i++ {
		if vet[i] < 0 {
			c := -vet[i]
			if c < 10 {
				mulhercalma = append(mulhercalma, vet[i])
			}
		}
	}
	return mulhercalma
}

func sortVet(vet []int) []int {

	for i := 0; i < len(vet); i++ {
		menor := i
		for j := i + 1; j < len(vet); j++ {
			if vet[j] < vet[menor] {
				menor = j
			}
		}
		vet[i], vet[menor] = vet[menor], vet[i]
	}

	return vet
}

func sortStress(vet []int) []int {
	for i := 0; i < len(vet); i++ {
		menor := i

		for j := i + 1; j < len(vet); j++ {
			x := vet[j]
			if x < 0 {
				x = -x
			}

			z := vet[menor]
			if z < 0 {
				z = -z
			} 
			if x < z {
				menor = j
			}
		}
		vet[i], vet[menor] = vet[menor], vet[i]
	}

	return vet
}

func reverse(vet []int) []int {
	reverso := []int{}

	for i := len(vet) - 1; i >= 0; i-- {
		reverso = append(reverso, vet[i])
	}
	return reverso
}

func unique(vet []int) []int {
	unico := []int{}

	for i := 0; i < len(vet); i++ {
		existe := false
		for j := 0; j < len(unico); j++ {
			 if vet[i] == unico[j] {
				existe = true
				break
			 }
			 
		}
		if !existe {
			unico = append(unico, vet[i])
		}
		
	}	
	return unico

}

func repeated(vet []int) []int {
	repetido := []int{}
	j := 0

	for i := 0; i < len(vet); i++ {
		repete := false
		for j = i; j >= 0; j-- {
			 if i != j && vet[i] == vet[j] {
				repete = true
				break
			 }
			 
		}
		if repete {
			repetido = append(repetido, vet[j])
		}
		
	}	
	return repetido

}

func main() {
	scanner := bufio.NewScanner(os.Stdin)

	for {
		if !scanner.Scan() {
			break
		}
		fmt.Print("$")
		line := scanner.Text()
		args := strings.Split(line, " ")
		fmt.Println(line)

		switch args[0] {
		case "get_men":
			printVec(getMen(str2vet(args[1])))
		case "get_calm_women":
			printVec(getCalmWomen(str2vet(args[1])))
		case "sort":
			printVec(sortVet(str2vet(args[1])))
		case "sort_stress":
			printVec(sortStress(str2vet(args[1])))
		case "reverse":
			array := str2vet(args[1])
			other := reverse(array)
			printVec(array)
			printVec(other)
		case "unique":
			printVec(unique(str2vet(args[1])))
		case "repeated":
			printVec(repeated(str2vet(args[1])))
		case "end":
			return
		}
	}
}

func printVec(vet []int) {
	fmt.Print("[")
	for i, val := range vet {
		if i > 0 {
			fmt.Print(", ")
		}
		fmt.Print(val)
	}
	fmt.Println("]")
}

func str2vet(s string) []int {
	if s == "[]" {
		return nil
	}
	s = s[1 : len(s)-1]
	parts := strings.Split(s, ",")
	var vet []int
	for _, part := range parts {
		n, _ := strconv.Atoi(part)
		vet = append(vet, n)
	}
	return vet
}

