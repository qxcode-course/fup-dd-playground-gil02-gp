package main
import "fmt"
func main() {
    var h, m, d, mes, ano int
    fmt.Scan(&h)
    fmt.Scan(&m)
    fmt.Scan(&d)
    fmt.Scan(&mes)
    fmt.Scan(&ano)

    fmt.Printf("%.2d:%.2d %.2d/%.2d/%.2d\n",h, m, d, mes, ano %100)
}
