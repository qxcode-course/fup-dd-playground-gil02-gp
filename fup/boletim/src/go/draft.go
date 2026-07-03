package main
import "fmt"
func main() {
    mat := make([][]int, 2)

    soma := 0
    for i := 0; i < 2; i++ {
        mat[i] = make([]int, 3)
        for j := 0; j < 3; j++ {
            fmt.Scan(&mat[i][j])
            soma += mat[i][j]
        }
    }
    fmt.Println(soma)
}