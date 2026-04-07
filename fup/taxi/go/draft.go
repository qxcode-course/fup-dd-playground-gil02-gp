package main
import "fmt"
func main() {
    var A, G, ra, rg float64
    fmt.Scan(&A)
    fmt.Scan(&G)
    fmt.Scan(&ra)
    fmt.Scan(&rg)

    alcool := A / ra
    gasolina := G / rg

    if alcool < gasolina {
        fmt.Println("A")
    } else {
        fmt.Println("G")
    }
}
