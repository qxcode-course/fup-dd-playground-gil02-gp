package main
import "fmt"
func main() {
    var n1, n2 int
    fmt.Scan(&n1, &n2)

    resultado := n1 / n2
    fmt.Printf("%d\n", resultado)
    fmt.Printf("%d\n", n1 % n2)
    var resultado2, n3, n4 float64
    n3 = float64(n1)
    n4 = float64(n2)
    resultado2 = n3 / n4
    fmt.Printf("%.2f\n", resultado2)
}