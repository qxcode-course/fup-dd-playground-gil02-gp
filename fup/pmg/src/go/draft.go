package main
import "fmt"
func media(vet []float64, qtd int) float64 {
    soma := 0.0
    for i := 0; i < qtd; i++ {
        soma += vet[i]
    }
    return soma / float64(qtd)
}
func main() {
    var qtd int 
    fmt.Scan(&qtd)
    var altura = make([]float64, qtd)
    for i := 0; i < qtd; i++ {
        fmt.Scan(&altura[i])
    }
    med := media(altura, qtd)
    fmt.Printf("%.2f\n", med)

    for i := 0; i < qtd; i++ {
        if altura[i] < med {
            fmt.Print("P")
        } else if altura[i] == med {
            fmt.Print("M")
        } else {
            fmt.Print("G")
        }
        if i < qtd-1 {
            fmt.Print(" ")
        }
    }
    fmt.Print("\n")
}
