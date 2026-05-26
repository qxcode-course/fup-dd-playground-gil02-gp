package main
import "fmt"
func consegueSair(p, s, e int) bool {
    posicao := 0
    for s > 0 {
        posicao += s
        if posicao >= p {
            return true
        }
        posicao -= e
        if posicao < 0 {
            return false
        }
        s -= 10
    }
    return false
}
func main() {
    var p, e int
    fmt.Scan(&p, &e)
    pulo := 1
    for {
        if consegueSair(p, pulo, e) {
            fmt.Println(pulo)
            break
        }
        pulo++
    }
}