package main
import "fmt"
func main() {
    var n1, n2, n3, n4 int
    fmt.Scan(&n1)
    fmt.Scan(&n2)
    fmt.Scan(&n3)
    fmt.Scan(&n4)
    jogos := n1 + n2 + n3 + n4

    if jogos == 0 {
        fmt.Println("nenhum")
    } else if jogos % 4 == 1 {
        fmt.Println("jog1")
    } else if jogos % 4 == 2 {
        fmt.Println("jog2")
    } else if jogos % 4 == 3 {
        fmt.Println("jog3")
    } else if jogos % 4 == 0 {
        fmt.Println("jog4")
    }
}
