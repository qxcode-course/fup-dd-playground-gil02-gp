package main
import "fmt"
func main() {
    var n int
    fmt.Scan(&n)
    if n <= 0{
        fmt.Printf("nao\n")
    } else {
        quadrado := 0
        for i := 0; i * i <= n; i++ {
            if i * i == n {
                quadrado = 1
                break
            }
        }
        if (quadrado == 1) {
            fmt.Printf("sim\n")
        } else {
            fmt.Printf("nao\n")
        }
    }

}
