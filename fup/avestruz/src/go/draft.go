package main
import (
    "fmt"
    "bufio"
    "os"
)
func main() {
    entrada := bufio.NewReader(os.Stdin)
    texto, _ := entrada.ReadString('\n')

    if len(texto) > 0 && texto[len(texto)-1] == '\n' {
        texto = texto[:len(texto)-1]
    }
    if len(texto) > 0 && texto[len(texto)-1] == '\r' {
        texto = texto[:len(texto)-1]
    }
    var letra string
    fmt.Fscan(entrada, &letra)
    alvo := letra[0]

    if alvo >= 'A' && alvo <= 'Z' {
        alvo += 'a' - 'A'
    }

    count := 0

    for i := 0; i < len(texto); i++ {
        c := texto[i]
        if c >= 'A' && c <= 'Z' {
            c += 'a' - 'A'
        }
        if c == alvo {
            count ++
        }
    }
    fmt.Println(count)
}