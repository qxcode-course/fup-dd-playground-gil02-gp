package main
import "fmt"
func main() {
    var n1, n2 int
    fmt.Scan(&n1, &n2)
    if n1 > n2 {
        fmt.Printf("%d\n", n1)
    } else {
        fmt.Printf("%d\n", n2)
    }
}
