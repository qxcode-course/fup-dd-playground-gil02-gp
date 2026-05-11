package main
import "fmt"
func main() {
    var n int
    fmt.Scan(&n)

    var calorias int
    soma := 0
    for i := 0; i < n; i++ {
        fmt.Scan(&calorias)
        soma += calorias
    }
    media := float64(soma) / float64(n)
    fmt.Printf("%.1f\n", media)
}
