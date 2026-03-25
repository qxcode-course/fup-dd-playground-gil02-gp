package main
import "fmt"
func main() {
    var hora, min, velocidade, consumo float64
    fmt.Scan(&min)
    fmt.Scan(&velocidade)
    fmt.Scan(&consumo)

    hora = min / 60.0
    distancia :=  velocidade * hora
    desempenho := distancia / consumo

    fmt.Printf("%.2f\n", desempenho)
}
