package main
import "fmt"
func main() {
    var n int
    fmt.Scan(&n)
    mat := make([][]int, n)

    for i := 0; i < n; i++ {
        mat[i] = make([]int, n)
        for j := 0; j < n; j++ {
            fmt.Scan(&mat[i][j])
        }
    }
    valorMaior := -1
    coluna := 0
    for j := 0; j < n; j++ {
        soma := 0
        for i := 0; i < n; i++ {
            soma += int(mat[i][j]) * int(mat[i][j])
        }
        if soma > valorMaior {
            valorMaior = soma
            coluna = j
        }
    }
    fmt.Printf("%d\n", coluna)
}