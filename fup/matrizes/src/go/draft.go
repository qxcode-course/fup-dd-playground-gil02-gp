package main
import "fmt"
func main() {
    var linhas, colunas int
    fmt.Scan(&linhas, &colunas)
    A := make([][]int, linhas)
    B := make([][]int, linhas)
    S := make([][]int, linhas)

    for i := 0; i < linhas; i++ {
        A[i] = make([]int, colunas)
        B[i] = make([]int, colunas)
        S[i] = make([]int, colunas)
    }

    for i := 0; i < linhas; i++ {
        for j := 0; j < colunas; j++ {
            fmt.Scan(&A[i][j])
        }
    }
    for i := 0; i < linhas; i++ {
        for j := 0; j < colunas; j++ {
            fmt.Scan(&B[i][j])
        }
    }
    for i := 0; i < linhas; i++ {
        for j := 0; j < colunas; j++ {
            S[i][j] = A[i][j] + B[i][j]
        }
    }
    for i := 0; i < linhas; i++ {
        fmt.Print("[ ")
        for j := 0; j < colunas; j++ {
            fmt.Print(S[i][j])
            if j < colunas-1 {
                fmt.Printf(" ")
            }
        }
        fmt.Print(" ]\n")
    }
}