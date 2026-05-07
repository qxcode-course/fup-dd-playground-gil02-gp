package main
import "fmt"
func main() {
    var p, s, e int
    fmt.Scan(&p, &s, &e)

    posicao := 0
    for {
        antes := posicao

        if s < 0 {
            s = 0
        }

        posicao += s

        if posicao >= p {
            fmt.Printf("%d saiu\n", antes)
            break
        } else {
            fmt.Printf("%d %d\n", antes, posicao)
        }

        posicao -= e

        if posicao < 0 {
            fmt.Printf("%d morreu\n", posicao)
            break
        }
        s -= 10
    
    }
}
