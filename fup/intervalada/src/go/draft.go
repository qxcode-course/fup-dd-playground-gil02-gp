package main
import "fmt"
func main() {
    var n, inferior, superior int
    fmt.Scan(&n)
    fmt.Scan(&inferior)
    fmt.Scan(&superior)

    contar := 0
    for i := 0; i < n; i++ {
        var num int
        fmt.Scan(&num)
        if num >= inferior && num <= superior {
            contar++
        }
    }
    fmt.Println(contar)
}
