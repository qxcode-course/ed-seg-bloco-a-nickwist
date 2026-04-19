package main
import "fmt"

func main() {
    var x, quant, quantgab int
    var faltam []int
    var gabarito []int
    deck := make(map[int]bool)

    fmt.Scan(&quantgab)
    fmt.Scan(&quant)

    gabarito = make([]int, 0)
    faltam = make([]int, 0)

    for i := 0; i < quant; i++ {
        fmt.Scan(&x)

        if deck[x] {
            gabarito = append(gabarito, x)
            deck[x] = true
        } else {
            deck[x] = true
        }
    }

    for i := 1; i <= quantgab; i++ {
        if !deck[i] {
            faltam = append(faltam, i)
        }
    }

    if len(gabarito) == 0 {
            fmt.Println("N")
    } else {
        for i := 0; i < len(gabarito); i++ {
            if i != len(gabarito) -1 {
                fmt.Printf("%d ", gabarito[i])
            } else {
                fmt.Println(gabarito[i])
            }
        }
    }
    
    if len(faltam) == 0 {
            fmt.Println("N")
    } else {
        for i := 0; i < len(faltam); i++ {

            if i != len(faltam) -1 {
                fmt.Printf("%d ", faltam[i])
            } else {
                fmt.Println(faltam[i])
            }
        } 
    }
}