package main
import "fmt"
func main() {
    var t int
    fmt.Scan(&t)
    var vet[50] int
    soma_jedi := 0
    soma_sith := 0

    for i := 0; i < t; i++ {
        fmt.Scan(&vet[i])
        if vet[i] < 1 || vet[i] > 10 {
            continue
        }
        if i < t/2 {
            soma_jedi += vet[i]
        } else {
            soma_sith += vet[i]
        }
    }

    if soma_jedi > soma_sith {
        fmt.Println("Jedi")
    } else if soma_sith > soma_jedi {
        fmt.Println("Sith")
    } else {
        fmt.Println("Empate")
    }

}
