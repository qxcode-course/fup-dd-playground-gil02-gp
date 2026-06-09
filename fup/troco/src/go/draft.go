package main
import "fmt"
func main() {
    var troco float64
    fmt.Scan(&troco)

    vet :=[]float64{100.00, 50.00, 20.00, 10.00, 5.00, 2.00, 1.00, 0.50, 0.25, 0.10, 0.05}
    
    cedulas := make([]int, len(vet))
    for troco > 0 && troco >= 0.05{
        for i := 0; i < len(vet); i++ {
            if troco >= vet[i] {
                cedulas[i] = (int)(troco / (vet[i]))
                troco = troco - (float64)(cedulas[i]) * (float64)(vet[i])
            } else {
                cedulas[i] = 0
            }
        }
    }
    for i := 0; i < len(cedulas); i++ {
        if cedulas[i] == 0 {
            continue
        }
        fmt.Printf("%d de %.2f\n", cedulas[i], vet[i])
    }
    if troco > 0 && troco < 0.05 {
        fmt.Printf("Falta %.2f\n", troco)
    }
}