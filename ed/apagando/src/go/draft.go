package main

import (
	"fmt"
)

func main() {
	var qntdpessoas, qntdsaida, pessoa int
	ids := []int{}
	idsnovo := []int{}
	sairam := make(map[int]bool)

	fmt.Scan(&qntdpessoas)
	for i := 0; i < qntdpessoas; i++ {
		fmt.Scan(&pessoa)
		ids = append(ids, pessoa)
	}

	fmt.Scan(&qntdsaida)

	for i := 0; i < qntdsaida; i++ {
		fmt.Scan(&pessoa)
		sairam[pessoa] = true
	}

	for _, pessoa := range ids {
		if !sairam[pessoa] {
			idsnovo = append(idsnovo, pessoa)	
		}
	}

	for _, pessoa := range idsnovo {
		fmt.Printf("%d ", pessoa)
	}

	fmt.Println("")
}
