package main

import (
    "bufio"
    "fmt"
    "os"
)

func main() {
    entrada := bufio.NewReader(os.Stdin)
    frase, _ := entrada.ReadString('\n')
    linha, _ := entrada.ReadString('\n')

    if len(frase) > 0 && frase[len(frase)-1] == '\n' {
        frase = frase[:len(frase)-1]
    }
    if len(frase) > 0 && frase[len(frase)-1] == '\r' {
        frase = frase[:len(frase)-1]
    }
    if len(linha) > 0 && linha[len(linha)-1] == '\n' {
        linha= linha[:len(linha)-1]
    }
    if len(linha) > 0 && linha[len(linha)-1] == '\r' {
        linha= linha[:len(linha)-1]
    }
    letra := linha[0]
    count := 0
    for i := 0; i < len(frase); i++ {
        if frase[i] == letra {
            count++
        }
    }
    fmt.Println(count)
}