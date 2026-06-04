package main

import "fmt"

func procurar_vivo(elementos []int, pos int) int {
    return pos % len(elementos)
}

func main() {

    var size, pos, f int
    fmt.Scan(&size, &pos, &f)
    pos--

    elementos := make([]int, size)

    if f == 1 {
        elementos[0] = 1
        for i := 1; i < size; i++ {
            if i % 2 == 0 {
                elementos[i] = i + 1
            } else {
                elementos[i] = -(i + 1)
            }
        }
    }

    for len(elementos) > 1 {

        pos = procurar_vivo(elementos, pos)

        fmt.Print("[ ")
        for i := 0; i < len(elementos); i++ {
            if i == pos && elementos[i] > 0 {
                fmt.Print(elementos[i], "> ")
            } else if i == pos && elementos [i] < 0 {
                fmt.Print("<", elementos[i], " ")
            } else {
                fmt.Print(elementos[i], " ")
            }
        }
        fmt.Println("]")

        var prox int

        if elementos[pos] > 0 {
            prox = (pos + 1) % len(elementos)
        } else {
            prox = (pos - 1 + len(elementos)) % len(elementos)
        }

        if elemento[pos] > 0 {
            elementos = append(elementos[:prox], elementos[prox+1:]...)
        } else {
            elementos = append(elementos[:prox], elementos[prox-1:]...)
        }


        pos = procurar_vivo(elementos, prox)
    }
    
    fmt.Printf("[ %d> ]\n", elementos[0])

}