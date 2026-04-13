package main
import "fmt"
func main() {
    var a, b int
    fmt.Scan(&a)
    fmt.Scan(&b)

    fmt.Printf("[ ")
    i := a
    for i < b {
        fmt.Printf("%d ", i)
        i++
    }
    fmt.Printf("]\n")
}
