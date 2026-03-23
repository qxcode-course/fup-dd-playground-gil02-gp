package main
import "fmt"
func main() {
    var tc float64
    fmt.Scan(&tc)

    tf := 1.8*tc + 32
    fmt.Printf("%.6f\n", tf)
}
