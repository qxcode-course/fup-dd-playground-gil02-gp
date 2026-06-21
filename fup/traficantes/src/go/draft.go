package main

import (
    "fmt"
	"bufio"
	"os"
	"strings"
)

func main() {
	texto := bufio.NewScanner(os.Stdin)

	texto.Scan()
    palavras := texto.Text()

	texto.Scan()
	procura := texto.Text()

	texto.Scan()
	trocar := texto.Text()

	palavras = strings.TrimSpace(palavras)
	procura = strings.TrimSpace(procura)
	trocar = strings.TrimSpace(trocar)

    for i := 0; i < len(palavras); i++ {
        if i+len(procura) <= len(palavras) && palavras[i:i+len(procura)] == procura {
            fmt.Print(trocar)
            i += len(procura) - 1
        } else {
            fmt.Printf("%c", palavras[i])
        }
    }
    fmt.Println()
}
