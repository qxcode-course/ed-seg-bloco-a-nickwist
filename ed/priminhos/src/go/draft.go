package main

import "fmt"

func pegarPrimo(primos []int, num, atual int) []int {

	if len(primos) <= num {
		for i := 2; i < atual; i++ {
			if atual%i == 0 {
				fmt.Println("a")
                break
				
			}
		}
	} else {
		return nil
	}

	primos = append(primos, atual)
	pegarPrimo(primos, num, atual+1)

	return primos
}

func main() {
	var num int
	var primos []int
	fmt.Scan(&num)

	primos = pegarPrimo(primos, num, 2)
	fmt.Println(primos)
}
