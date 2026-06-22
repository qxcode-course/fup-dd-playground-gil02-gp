package main
import (
    "fmt"
    "bufio"
    "os"
)
func main() {
    entrada := bufio.NewReader(os.Stdin)
    nome, _ := entrada.ReadString('\n')
    if len(nome) > 0 && nome[len(nome)-1] == '\n' {
        nome = nome[:len(nome)-1]
    }
    soma := 0
    for i := 0; i < len(nome); i++{
        soma += int(nome[i])
    }
    fmt.Println(soma % 50)
}