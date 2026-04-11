package main
import "fmt"
func main() {
    var n, a, d int
    fmt.Scan(&n)
    fmt.Scan(&d)
    fmt.Scan(&a)
    
    movimento := (d - a + n) % n
    
    fmt.Println(movimento)
}
