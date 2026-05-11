package main
import "fmt"
func main() {
    var n int
    n = 5

    vetor := make([]int, n)
    for i := 0; i < n; i++ {
        fmt.Scan(&vetor[i])
    }
    menor := vetor[0]
    for i := 0; i < n; i++ {
        if vetor[i] < menor {
            menor = vetor[i]
        }
    }
    fmt.Printf("%d\n", menor)   
}
