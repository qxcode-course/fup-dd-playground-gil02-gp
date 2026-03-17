package main
import "fmt"
func main() {
    var a, b, c int
    fmt.Scan(&a)
    fmt.Scan(&b)
    fmt.Scan(&c)
    var p1, p2, p3, dinheiro float64
    fmt.Scan(&p1)
    fmt.Scan(&p2)
    fmt.Scan(&p3)
    fmt.Scan(&dinheiro)

    total := float64(a)*p1 + float64(b)*p2 + float64(c)*p3
    troco := dinheiro - total
    fmt.Printf("%.2f\n", troco)
}