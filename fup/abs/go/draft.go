package main

import (
	"fmt"
	"math"
)
func main() {
    var a, b float64
    fmt.Scan(&a, &b)

    valor := math.Abs(a - b)
    fmt.Printf("%.f\n", valor)
}
