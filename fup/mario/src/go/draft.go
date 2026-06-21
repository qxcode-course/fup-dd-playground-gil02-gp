package main
import "fmt"
func main() {
    var n int
    fmt.Scan(&n)

    vetor := make([]int, n)
    maior := 0
    for i := 0; i < n; i++ {
        fmt.Scan(&vetor[i])
        if vetor[i] > maior {
            maior = vetor[i]
        }
    }
    for nivel := maior; nivel > 0; nivel -- {
        for i := 0; i < n; i++ {
            if vetor[i] >= nivel {
                fmt.Print("#")
            } else {
                fmt.Print("_")
            }
        }
        fmt.Println()
    }
}