package main
import "fmt"
func main() {
    var H, P, F, D int

    fmt.Scan(&H, &P, &F, &D)

    if D == 1 {
        for F != H {

            if F == 16 {
                F = 0
                continue
            }
            
            if F == P {
                fmt.Println("N")
                return
            }

            F++
        }
        fmt.Println("S")
    } else if D == -1 {
        for F != H {
            
            if F == P {
                fmt.Println("N")
                return
            }

            if F == 0 {
                F = 15
                continue
            }

            F--
        }
        fmt.Println("S")
    } else {}
}
