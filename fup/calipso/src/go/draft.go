package main
import (
    "fmt"
    "bufio"
    "os"
)
func main() {
    scanner := bufio.NewReader(os.Stdin)

    var n int
    fmt.Fscanf(scanner, "%d", &n)
    scanner.ReadByte()

    for ; n > 0; n-- {
        frase, _ := scanner.ReadString('\n')
        if len(frase) > 0 && frase[len(frase)-1] == '\n' {
            frase = frase[:len(frase)-1]
        }
        maiuscula := false
        for _, c := range frase {
            if c != ' ' {
                if c >= 'A'&& c <= 'Z' {
                    maiuscula = true
                }
                break
            }
        }
        for _, c := range frase {
            if c == ' ' {
                fmt.Print(" ")
                continue
            }
            if maiuscula {
                if c >= 'a' && c <= 'z' {
                    c -= 'a' - 'A'
                }
            } else {
                if c >= 'A' && c <= 'Z' {
                    c += 'a' - 'A'
                }
            }
            fmt.Printf("%c", c)
            maiuscula = !maiuscula
        }
        fmt.Println()
    }
}