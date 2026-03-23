package main
import "fmt"
import "math"
func main() {
    var a, b, c, p float64
    fmt.Scan(&a, &b, &c)

    p = (a + b + c) / 2.0

    area := math.Sqrt(p * (p - a) * (p - b) * (p - c))
    fmt.Printf("%.2f\n", area)
}
