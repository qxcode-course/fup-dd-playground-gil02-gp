package main
import (
    "fmt"
    "bufio"
    "os"
)
func main() {
    frase := bufio.NewReader(os.Stdin)
    texto, _ := frase.ReadString('\n')

    for _, char := range texto {
        if char >= 'a' && char <= 'z' {
            fmt.Printf("%c", char - ('a' - 'A'))
        } else if char >= 'A' && char <= 'Z' {
            fmt.Printf("%c", char + ('a' - 'A'))
        } else {
            fmt.Printf("%c", char)
        }
    }

}