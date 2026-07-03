package main
import "fmt"
func main() {
    var mat[3][3] int
    for i := 0; i < 3; i++ {
        for j := 0; j < 3; j++ {
            fmt.Scan(&mat[i][j])
        }
    }
    simetrica := 1
    for i := 0; i < 3; i++ {
        for j := 0; j < 3; j++ {
            if mat[i][j] != mat[j][i] {
                simetrica = 0
            }
        }
    }
    if simetrica == 1 {
        fmt.Print("sim\n")
    } else {
        fmt.Print("nao\n")
    }
}