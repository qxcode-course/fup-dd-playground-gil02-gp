package main
import "fmt"

func absDiff(a, b int) int {
    if a > b {
        return a - b
    }
    return b - a
}

func main() {
    var chute_chico int
    var chute_cebolinha int
    var qtd_animais int

    fmt.Scan(&chute_chico)
    fmt.Scan(&chute_cebolinha)
    fmt.Scan(&qtd_animais)

    total := 0
    for i := 0; i < qtd_animais; i++ {
        var animal string
        fmt.Scan(&animal)
        switch animal {
        case "v", "c":
            total += 4
        case "g":
            total += 2
        }
    }
    fmt.Println(total)
    
    chute_chico = absDiff(chute_chico, total)
    chute_cebolinha = absDiff(chute_cebolinha, total)

    if chute_chico < chute_cebolinha {
        fmt.Println("Chico Bento")
    } else if chute_cebolinha < chute_chico {
        fmt.Println("Cebolinha")
    } else {
        fmt.Println("empate")
    }
}
