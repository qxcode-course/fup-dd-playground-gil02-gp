package main
import (
    "fmt"
    "bufio"
    "os"
)
func main() {
    scanner := bufio.NewScanner(os.Stdin)
    if scanner.Scan() {
        frase := scanner.Text()
    vogais := ""
    consoantes := ""

    for _, c := range frase {
        if c == ' ' || c == '\n' {
            continue
        }
        if c == 'a' || c == 'e' || c == 'i' || c == 'o' || c == 'u' {
            vogais += string(c)
        } else {
            consoantes += string(c)
        }
    }
    fmt.Println(vogais)
    fmt.Println(consoantes)
    }
}