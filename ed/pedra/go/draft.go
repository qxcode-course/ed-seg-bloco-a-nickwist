package main
import "fmt"
func main() {
    var NC, i, ganhador, arrem1, arrem2, pontuacao, pontuacaoganhador int

    fmt.Scan(&NC)

    for i = 0; i<NC; i++ {
        fmt.Scan(&arrem1, &arrem2)
        if arrem1 < 10 || arrem2 < 10 {
            continue
        } 
        
        if arrem1 < arrem2 {
            pontuacao = arrem2 - arrem1
        } else {
            pontuacao = arrem1 - arrem2
        }

        if pontuacaoganhador == 0 {
            ganhador = i
            pontuacaoganhador = pontuacao
        } else if pontuacaoganhador > pontuacao {
            ganhador = i
            pontuacaoganhador = pontuacao
        } else {}
    }
    if ganhador == 0 {
        fmt.Println("sem ganhador")
        return
    }
    fmt.Printf("%d\n", ganhador)
}
