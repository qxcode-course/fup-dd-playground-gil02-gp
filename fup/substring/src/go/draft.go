package main
import (
    "bufio"
    "fmt"
    "os"
)
func main() {
    palavra := bufio.NewReader(os.Stdin)
    var texto string
    texto, _ = palavra.ReadString('\n')
    if len(texto) > 0 && texto[len(texto)-1] == '\n' {
        texto = texto[:len(texto)-1]
    }
    var inicio, qtd int
    fmt.Fscan(palavra, &inicio)
    fmt.Fscan(palavra, &qtd)

    if inicio < 0 || inicio >= len(texto) || qtd < 0 {
        fmt.Println("")
        return
    }

    fim := inicio + qtd
    if fim > len(texto) {
        fim = len(texto)
    }
    fmt.Println(texto[inicio:fim])
}