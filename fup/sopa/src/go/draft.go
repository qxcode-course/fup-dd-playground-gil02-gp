package main
import "fmt"
func main() {
    var n int
    fmt.Scan(&n)

    if n == 0 || n == 1 {
        fmt.Println(1)
        return
    }
    
    var a, b int = 1, 1
    for i := 2; i < n; i++ {
        a, b = b, a + b
    }
    fmt.Println(b)
}