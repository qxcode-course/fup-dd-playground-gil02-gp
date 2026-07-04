package main
import "fmt"
func main() {
    var mat[5][5] int
    somaPrincipal := 0
    somaSecundario := 0
    for i := 0; i < 5; i++ {
        for j := 0; j < 5; j++ {
            fmt.Scan(&mat[i][j])
        }
    }
    for i := 0; i < 5; i++ {
        somaPrincipal += mat[i][i]
        somaSecundario += mat[i][4-i]
    }
    resul := somaPrincipal - somaSecundario
    fmt.Println(resul)
}