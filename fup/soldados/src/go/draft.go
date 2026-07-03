package main
import "fmt"
func main() {
    var nl, nc int
    fmt.Scan(&nl, &nc)

    mat := make([][]int, nl)
    for i := 0; i < nl; i++ {
        mat[i] = make([]int, nc)
        for j := 0; j < nc; j++ {
            fmt.Scan(&mat[i][j])
        }
    }
    count := 0
    for j := 0; j < nc; j++ {
        for i := 0; i < nl-1; i++ {
            if mat[i][j] > mat[i+1][j] {
                count++
            }
        }
    }
    fmt.Println(count)
}