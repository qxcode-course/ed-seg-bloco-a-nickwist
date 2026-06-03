package main
import "fmt"
func main() {
    var busca, consulta []string
    var res []int
    var quantbusca, quantconsulta int

    fmt.Scan(&quantbusca)

    for i := 0; i < quantbusca; i++ {
        var buscas string
        fmt.Scan(&buscas)
        busca = append(busca, buscas)
    }

    fmt.Scan(&quantconsulta) 

    for i := 0; i < quantconsulta; i++ {
        var consultas string
        fmt.Scan(&consultas)
        consulta = append(consulta, consultas)
    }

    for i := 0; i < len(consulta); i++ {
        quantidade := 0
        for j := 0; j < len(busca); j++ {
            if consulta[i] == busca[j] {
                quantidade++
            }
        }
        res = append(res, quantidade)
    }

    for i := 0; i < len(res); i++ {
        if i != len(res)- 1 {
            fmt.Printf("%d ", res[i])
        } else {
            fmt.Println(res[i])
        }
    }
}