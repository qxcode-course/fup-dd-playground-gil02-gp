package main
import "fmt"
func main() {
    var a, b int
    fmt.Scan(&a, &b)

    fim := b
    fmt.Print("[ ")
    for i := a; i <= fim; i++ {
        fmt.Print(a, " ", b, " ")
        a++
        b--
    }
    fmt.Print("]\n")
}
