package main
import "fmt"
import "math"
func main() {
    var a, b, c, x, num float64
    fmt.Scan(&a)
    fmt.Scan(&b)
    fmt.Scan(&c)

    num = math.Pow(b,2) - 4 * a * c

    x = (b*(-1) + math.Sqrt(num)) / (2 * a)
    x1 := (b*(-1) - math.Sqrt(num)) / (2 * a)
    

    if num > 0{
        fmt.Printf("%.2f\n%.2f\n", x, x1)
    } else if num == 0{
        fmt.Printf("%.2f\n", x)
    } else {
        fmt.Printf("nao ha raiz real\n")
    }
}
