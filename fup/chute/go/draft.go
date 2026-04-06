package main
import "fmt"
import "math"
func main() {
    var valor_inteiro int
    fmt.Scan(&valor_inteiro)
    var primeiro_chute, segundo_chute int
    fmt.Scan(&primeiro_chute)
    fmt.Scan(&segundo_chute)
    
    valor1 := math.Abs(float64(valor_inteiro - primeiro_chute))
    valor2 := math.Abs(float64(valor_inteiro - segundo_chute))
    if valor1 < valor2 {
        fmt.Printf("primeiro\n")
    } else if valor2 < valor1 {
        fmt.Printf("segundo\n")
    } else {
        fmt.Printf("empate\n")
    }
}
