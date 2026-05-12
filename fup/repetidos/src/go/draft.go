package main
import "fmt"
func main() {
    var p, n int
    fmt.Scan(&p, &n)
    contador := 0
    var vetor []int
    vetor = make([]int, n)
    for i := 0; i < n; i++{
        fmt.Scan(&vetor[i])
    }
    for i := 0; i < n; i++{
        if vetor[i] == p {
            contador++
        }
    }
    fmt.Println(contador)
}
