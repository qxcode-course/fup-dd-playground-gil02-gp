package main
import "fmt"
func main() {
    var n, x int
    var machos [51]int
    var femeas [51]int
    fmt.Scan(&n)

    casal := 0
    for i := 0; i < n; i++ {
        fmt.Scan(&x)
        especiesID := x
        if especiesID < 0 {
            especiesID = -especiesID
            femeas[especiesID]++
        } else {
            machos[especiesID]++
        }
    }

    for i := 1; i <= 50; i++ {
        if machos[i] < femeas[i] {
            casal += machos[i]
        } else {
            casal += femeas[i]
        }
    }

    fmt.Println(casal)
}
