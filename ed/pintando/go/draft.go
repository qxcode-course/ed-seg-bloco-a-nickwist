package main
import ("fmt"
        "math"
        )
func main() {
    var n1, n2, n3, semip, area float64

    fmt.Scan(&n1, &n2, &n3)

    semip = (n1 + n2 + n3)/2
    area = math.Sqrt(semip*(semip-n1)*(semip-n2)*(semip-n3))

    fmt.Printf("%.2f\n", area)

}
