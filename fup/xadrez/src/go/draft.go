package main
import "fmt"
func main() {
    var l, c int
    fmt.Scan(&l, &c)
    tabu := l + c
    if tabu % 2 == 0 {
        fmt.Println("1")
    } else {
        fmt.Println("0")
    }
}
