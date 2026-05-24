package main
import "fmt"
func abs(x float64) float64 {
    if x < 0 {
        return -x
    }
    return x
}
func main() {
    var n int
    fmt.Scan(&n)

    valores := make([]float64, n)
    chutes := make([]float64, n)
    escolha := make([]string, n)

    for i := 0; i < n; i++ {
        fmt.Scan(&valores[i])
    }
    for i := 0; i < n; i++ {
        fmt.Scan(&chutes[i])
    }
    for i := 0; i < n; i++ {
        fmt.Scan(&escolha[i])
    }
    p1 := 0
    p2 := 0

    for i := 0; i < n; i++ {
        if abs(valores[i]-chutes[i]) < 0.000001 {
            p1++
        } else if escolha[i] == "M" {
            if valores[i] > chutes[i] {
                p2++
            } else {
                p1++
                }
            } else {
                if valores[i] < chutes[i] {
                    p2++
                } else {
                    p1++
                }
            }
        }

    if p1 > p2 {
        fmt.Println("primeiro")
    } else if p2 > p1 {
        fmt.Println("segundo")
    } else {
        fmt.Println("empate")
    }

}
