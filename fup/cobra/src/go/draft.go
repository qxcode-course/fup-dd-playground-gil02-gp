package main
import "fmt"
func main() {
    var n, x, y, s int
    var c string
    fmt.Scan(&n)
    fmt.Scan(&x)
    fmt.Scan(&y)
    fmt.Scan(&c)
    fmt.Scan(&s)

    if c == "R" {
     x = (x + s) % n
    } else if c == "L" {
     x = (x - s%n + n) % n
    } else if c == "D" {
        y = (y + s) % n
    } else if c == "U" {
        y = (y - s%n + n) % n
    }
    fmt.Println(x, y)
}
