package main
import (
    "fmt"
    "bufio"
    "os"
    "strings"
)
func main() {
    texto := bufio.NewReader(os.Stdin)
    frase, _ := texto.ReadString('\n')
    trecho, _ := texto.ReadString('\n')

    frase = strings.TrimSpace(frase)
    trecho = strings.TrimSpace(trecho)

    fmt.Println(strings.Count(frase, trecho))
}