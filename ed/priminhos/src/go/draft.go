package main

import "fmt"

func pegarPrimo(primos []int, num, atual int) []int {
	if len(primos) >= num {
		return primos
	}

	ehPrimo := true

	for i := 2; i < atual; i++ {
		if atual%i == 0 {
			ehPrimo = false
			break
		}
	}

	if ehPrimo {
		primos = append(primos, atual)
	}

	return pegarPrimo(primos, num, atual+1)
}

func main() {
	var num int
	var primos []int
	fmt.Scan(&num)

	primos = pegarPrimo(primos, num, 2)
	fmt.Print("[")
	for i := 0; i < len(primos); i++ {
		if i != len(primos)-1 {
			fmt.Printf("%d, ", primos[i])
		} else {
			fmt.Printf("%d", primos[i])
		}
	}
	fmt.Println("]")
}
