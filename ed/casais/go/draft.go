package main
import "fmt"

func main() {
    var pares, i, quant, x int
    animal := make(map[int]bool)

    fmt.Scan(&quant)

    for i = 0; i < quant; i++ {
        fmt.Scan(&x)

        if animal[-x] {
            pares++
            delete(animal, -x)
        } else {
            animal[x] = true
        }
    }

    fmt.Println(pares)
}