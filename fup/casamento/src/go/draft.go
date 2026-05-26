package main
import "fmt"
func main() {
    var vet[5] int
    for i := 0; i < 5; i++ {
        fmt.Scan(&vet[i])
    }
    maior := vet[0]
    menor := vet[0]
    for i := 1; i < 5; i++ {
        if vet[i] > maior {
            maior = vet[i]
        }
        if vet[i] < menor {
            menor = vet[i]
        }
    }
    fmt.Println(maior + menor)
}