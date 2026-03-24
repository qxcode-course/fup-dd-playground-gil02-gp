package main

import (
	"fmt"
	"math"
)
func main() {
    var a, b float64
    fmt.Scan(&a, &b)
    maior := math.Max(a, b)
    fmt.Printf("%.f\n", maior)
}
