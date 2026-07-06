package main
import "fmt"

func somar(nums []int, valor, inicio, soma int) bool {

    if soma == valor {
        return true
    }

    for i := inicio; i < len(nums); i++ {
        soma += nums[i]

        if somar(nums, valor, i+1, soma) {
            return true
        }

        soma -= nums[i]
    }

    return false
}

func main() {
    var quant, valor, numero int
    var nums []int

    fmt.Scan(&quant, &valor)

    for i := 0; i < quant; i++ {
        fmt.Scan(&numero)
        nums = append(nums, numero)
    }
    
    fmt.Println(somar(nums, valor, 0, 0))
}