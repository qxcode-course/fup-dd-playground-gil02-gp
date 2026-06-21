package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)
func main() {
    scanner := bufio.NewScanner(os.Stdin)
    var palavras []string
    if scanner.Scan() {
        texto := strings.TrimSpace(scanner.Text())
        palavras = strings.Fields(texto)
    }
    if err := scanner.Err(); err != nil {
        fmt.Println("Erro ao ler a entrada:", err)
        return
    }

    ordenadas := true
    for i := 0; i < len(palavras)-1; i++ {
        if palavras[i] > palavras[i+1] {
            ordenadas = false
            break
        }
    }
    if ordenadas {
        fmt.Println("sim")
    } else {
        fmt.Println("nao")
    }
}