package main
import "fmt"
func main() {
    var n int
    fmt.Scan(&n)
    
    vetor := make([]int, n)
    for i := 0; i < n; i++ {
        fmt.Scan(&vetor[i])
    }
    if n == 0 {
            fmt.Print("\n")
        }
    
    for i := 0; i < n; i++ {
        fmt.Printf("%d\n", vetor[i])
    }
}
