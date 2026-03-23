package main
import "fmt"
import "math"
func main() {
    var x1, x2, y1, y2, distancia float64
    fmt.Scan(&x1, &y1)
    fmt.Scan(&x2, &y2)
    
    distancia = math.Sqrt((x2 - x1)*(x2 - x1) + (y2 - y1)*(y2 - y1))
    fmt.Printf("%.2f\n", distancia)
}
