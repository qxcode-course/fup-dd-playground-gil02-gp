package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)
func main() {
    scanner := bufio.NewScanner(os.Stdin)
    scanner.Scan()
    frase := scanner.Text()
    frase = strings.TrimSpace(frase)

    elementos := strings.Fields(frase)
    for i, e := range elementos {
        temLetra := false
        temPonto := false
        for _, c := range e {
            if c >= 'a' && c <= 'z' || c >= 'A' && c <= 'Z' {
                temLetra = true
                break
            }
            if c == '.' {
                temPonto = true
            }
        }
        if temLetra {
            fmt.Print("str")
        } else if temPonto {
            fmt.Print("float")
        } else {
            fmt.Print("int")
        }
        if i < len(elementos)-1 {
            fmt.Print(" ")
        }
    }
    fmt.Println()
}