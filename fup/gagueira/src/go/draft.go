package main
import (
    "fmt"
    "bufio"
    "os"
)
func main() {
    entrada := bufio.NewReader(os.Stdin)
    frase, _ := entrada.ReadString('\n')

    if len(frase) > 0 && frase[len(frase)-1] == '\n' {
        frase = frase[:len(frase)-1]
    }
    palavra := ""
    for i := 0; i <= len(frase); i++ {
        if i == len(frase) || frase[i] == ' ' {
            fmt.Print(palavra, " ", palavra)
            if i != len(frase){
                fmt.Print(" ")
            }
            palavra = ""
        } else {
            palavra += string(frase[i])
        }
    }
    fmt.Println()
}