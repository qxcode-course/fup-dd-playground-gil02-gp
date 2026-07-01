package main
import "fmt"
func main() {
    var mat[3][3] int
    for i := 0; i < 3; i++ {
        for j := 0; j < 3; j++ {
            fmt.Scan(&mat[i][j])
        }
    }
    soma := 0
    for j := 0; j < 3; j++ {
        soma += mat[0][j]
    }
    magico := 1
    for i := 1; i < 3; i++ {
        linha := 0
        for j := 0; j < 3; j++ {
            linha += mat[i][j]
        }
        if linha != soma {
            magico = 0
            break
        }
    }
    if magico == 1 {
        for j := 0; j < 3; j++ {
            coluna := 0
            for i := 0; i < 3; i++ {
                coluna += mat[i][j]
            }
            if coluna != soma {
                magico =  0
                break
            }
        }
    }
    if magico == 1{
        diagonal1 := 0
        for i := 0; i < 3; i++ {
            diagonal1 += mat[i][i]
        }
        if diagonal1 != soma {
            magico = 0
        }
    }
    if magico == 1{
        diagonal2 := 0
        for i := 0; i < 3; i++ {
            diagonal2 += mat[i][2-i]
        }
        if diagonal2 != soma {
            magico = 0
        }
    }
    if magico == 1 {
        fmt.Printf("sim\n")
    } else {
        fmt.Printf("nao\n")
    }
}