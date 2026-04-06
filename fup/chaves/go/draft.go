package main
import "fmt"
func main() {
    var a int
    fmt.Scan(&a)

    if a > 1 {
        fmt.Println("positivo")
    } else if a <= -1 {
        fmt.Println("negativo")
    } else {
        fmt.Println("nulo")
    }
}
