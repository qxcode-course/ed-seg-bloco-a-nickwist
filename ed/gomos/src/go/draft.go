package main
import "fmt"
func main() {
    var quantG int
    var mov string
    var x []int
    var y []int

    x = make([]int, 0)
    y = make([]int, 0)

    fmt.Scan(&quantG, &mov)

    for i := 0; i < quantG; i++ {
        var z, w int

        fmt.Scan(&z, &w)
        x = append(x, z)
        y = append(y, w)
    }

    if mov == "L" {
        for i := quantG-1; i >= 0; i-- {
            if i != 0 {
                x[i] = x[i-1]
                y[i] = y[i-1]
            } else {
                x[i]--
            }
        }
    }

    if mov == "R" {
        for i := quantG-1; i >= 0; i-- {
            if i != 0 {
                x[i] = x[i-1]
                y[i] = y[i-1]
            } else {
                x[i]++
            }
        }
    }

    if mov == "U" {
        for i := quantG-1; i >= 0; i-- {
            if i != 0 {
                y[i] = y[i-1]
                x[i] = x[i-1]
            } else {
                y[i]--
            }
        }
    }

    if mov == "D" {
        for i := quantG-1; i >= 0; i-- {
            if i != 0 {
                y[i] = y[i-1]
                x[i] = x[i-1]
            } else {
                y[i]++
            }
        }
    }
    
    for i := 0; i < len(x); i++ {
            fmt.Println(x[i], y[i])
    }
}