package main
import "fmt"
func main() {
    var n int
    fmt.Scan(&n)
    var aluno[50] int
    var servidores[50] int
    cont_alu := 0
    cont_ser := 0
 
    for i := 0; i < n; i++ {
        var tipo int
        fmt.Scan(&tipo)
        if tipo % 2 == 1 {
            aluno[cont_alu] = tipo
            cont_alu++
        } else {
            servidores[cont_ser] = tipo
            cont_ser++
        }
    }
    fmt.Print("[")
    for i := 0; i < cont_alu; i++ {
        fmt.Print(" ", aluno[i])
    }
    fmt.Print(" ]\n")

    fmt.Print("[")
    for i := 0; i < cont_ser; i++ {
        fmt.Print(" ", servidores[i])
    }
    fmt.Print(" ]\n")
}
