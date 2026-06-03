package main
import "fmt"
func main() {
    var tamanho, qntdRotacoes int
    var vetor, res []int

    fmt.Scan(&tamanho, &qntdRotacoes)

    for i := 0; i < tamanho; i++ {
        var v int
        fmt.Scan(&v)

        vetor = append(vetor, v)
    }
    
    if tamanho != 0 {
        qntdRotacoes = tamanho - (qntdRotacoes % tamanho)
        }

    for i := 0; i < tamanho; i++ {
        res = append(res, vetor[(i+qntdRotacoes)%tamanho])
    }

    fmt.Print("[ ")
    for i := 0; i < tamanho; i++ {
        fmt.Printf("%d ", res[i])
    }
    fmt.Println("]")
}