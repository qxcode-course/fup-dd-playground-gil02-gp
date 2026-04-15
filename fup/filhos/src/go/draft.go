package main
import "fmt"
func main() {
    var filhos, n int
    fmt.Scan(&filhos, &n)
    for i := 0; i < n; i++ {
        fmt.Println(filhos + i*2)
    }
    
}
