package main

import "fmt"

func procurar_vivo(elementos []int, pos int) int {
    return pos % len(elementos)
}

func main() {

    var size, pos int
    fmt.Scan(&size, &pos)
    pos--

    elementos := make([]int, size)

    for i := 0; i < size; i++ {
        elementos[i] = i + 1
    }

    for len(elementos) > 1 {

        pos = procurar_vivo(elementos, pos)

        fmt.Print("[ ")
        for i := 0; i < len(elementos); i++ {
            if i == pos {
                fmt.Print(elementos[i], "> ")
            } else {
                fmt.Print(elementos[i], " ")
            }
        }
        fmt.Println("]")

        prox := (pos + 1) % len(elementos)

        elementos = append(elementos[:prox], elementos[prox+1:]...)

        pos = procurar_vivo(elementos, prox)
    }
    
    fmt.Printf("[ %d> ]\n", elementos[0])

}