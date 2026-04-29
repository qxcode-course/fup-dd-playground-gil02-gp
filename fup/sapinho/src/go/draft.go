package main
import "fmt"
func main() {
    var p, s, e int
    fmt.Scan(&p, &s, &e)

    posicao := 0
    for {
        nova := posicao + s
        if nova >= p {
            fmt.Println(posicao, "saiu")
            break
        }
        fmt.Println(posicao, nova)
    posicao = nova - e
    
    }
}
