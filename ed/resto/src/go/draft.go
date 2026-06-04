package main

import "fmt"

func div(numero int) {
	if numero == 1 {
		fmt.Println("0 1")
		return
	}

	r := numero % 2

	div(numero / 2)

	fmt.Println(numero/2, r)

}

func main() {
	var numero int

	fmt.Scan(&numero)

	div(numero)

}
