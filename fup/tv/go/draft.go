package main
import "fmt"
func main() {
    var valor float64
    var parcelas int
    fmt.Scan(&valor)
    fmt.Scan(&parcelas)

    juro := float64(parcelas-1) * 0.05
    total := valor * (1 + juro)
    parcela := total / float64(parcelas)

    fmt.Printf("%.2f\n", parcela)
    fmt.Printf("%.2f\n", total)
}
