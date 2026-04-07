package main
import "fmt"

func main() {
    var chute, real float64
    var escolha string

    fmt.Scanln(&chute)
    fmt.Scanln(&escolha)
    fmt.Scanln(&real)

    if chute == real {
        fmt.Println("primeiro")
    } else if (escolha == "M" && real > chute) || (escolha == "m" && real < chute) {
        fmt.Println("segundo")
    } else {
        fmt.Println("primeiro")
    }
}
