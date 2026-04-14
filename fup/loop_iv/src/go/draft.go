package main
import "fmt"
func main() {
    var a, b int
    fmt.Scan(&a, &b)
    fmt.Printf("[ ")
    if b != a {
        for i := a; i < b; i++ {
            fmt.Printf("%d ", i)
        }
    }
    fmt.Printf("]\n")
}
