package main
import "fmt"
func main() {
    var l, r, d int
    fmt.Scan(&l, &r, &d)
    if r > 50 && l < r && r > d {
        fmt.Printf("S\n")
    } else {
        fmt.Printf("N\n")
    }
}
