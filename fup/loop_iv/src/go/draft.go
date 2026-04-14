package main
import "fmt"
func main() {
    var a, b int
    fmt.Scan(&a, &b)
    fmt.Printf("[ ")
    step := 1
    if a > b {
        step = -1
    }
    if a != b {
        for i := a; i != b; i += step {
            fmt.Printf("%d ", i)
        }
    }
    fmt.Printf("]\n")
}
