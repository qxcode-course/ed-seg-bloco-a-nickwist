package main
import "fmt"
func main() {
    var pares, i, quant, x int
    var animal map[int] bool
    animal = make(map[int] bool)

    fmt.Scan(&quant)

    for i = 0; i < quant; i++ {
        fmt.Scan(&x)
        animal[x] = true
        if animal[-x] {
            pares++
            delete(animal, -x)
            delete(animal, x)
        }

        
    }
    fmt.Println(pares)

    
}
