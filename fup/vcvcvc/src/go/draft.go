package main
import (
    "fmt"
    "bufio"
    "os"
)
func main() {
    scanner := bufio.NewScanner(os.Stdin)
    scanner.Scan()
    frase := scanner.Text()

    for _, vogal := range frase {
        if vogal == '\n' {
            break
        }
        if vogal == ' ' {
            fmt.Print(" ")
        } else if vogal == 'a' || vogal == 'e' || vogal == 'i' || vogal == 'o' || vogal == 'u' ||
        vogal == 'A' || vogal == 'E' || vogal == 'I' || vogal == 'O' || vogal == 'U' {
            fmt.Print("v")
        } else {
            fmt.Print("c")
        }
    }
    fmt.Println()
}