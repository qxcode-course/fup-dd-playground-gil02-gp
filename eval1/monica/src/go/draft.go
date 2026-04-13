package main
import "fmt"
func main() {
    var m, a, b int
    fmt.Scan(&m, &a, &b)

    c := m - (a + b)
    mais_velhos := a
    if b > a && b > c {
        mais_velhos = b
    } else if c > a && c > b {
        mais_velhos = c
    }
    fmt.Printf("%d\n", mais_velhos)
    
}
