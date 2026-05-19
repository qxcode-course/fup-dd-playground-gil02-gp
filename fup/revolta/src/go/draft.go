package main
import "fmt"
func main() {
    var n, x int
    fmt.Scan(&n)
    soma_pares := 0
    soma_impares := 0

    for i := 0; i < n; i++ {
        fmt.Scan(&x)

        if x < 1 || x > 50{
            continue
        }
        if x%2 == 0 {
            soma_pares += x
        }   else {
            soma_impares += x
        }
    }
    if soma_impares > soma_pares {
        fmt.Printf("soldados\n")
    } else if soma_pares > soma_impares{
        fmt.Printf("rebeldes\n")
    } else {
        fmt.Printf("empate\n")
    }
}
