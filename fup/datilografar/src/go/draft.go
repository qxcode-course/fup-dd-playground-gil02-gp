package main
import (
    "fmt"
    "bufio"
    "os"
)
func main() {
    entrada := bufio.NewReader(os.Stdin)
    var quebrado string
    var numero string

    fmt.Fscan(entrada, &quebrado, &numero)

    resul := ""

    for i := 0; i < len(numero); i++{
        if numero[i] != quebrado[0] {
            resul += string(numero[i])
        }
    }
    i := 0
    for i < len(resul) && resul[i] == '0' {
        i++
    }
    if i == len(resul){
        fmt.Println(0)
    } else {
        fmt.Println(resul[i:])
    }
}