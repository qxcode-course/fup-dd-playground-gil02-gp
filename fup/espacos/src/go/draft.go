package main
import (
    "fmt"
    "bufio"
    "os"
)
func main() {
    scanner := bufio.NewScanner(os.Stdin)
    scanner.Scan()
    texto := scanner.Text()

    espaco := false

    for _, c := range texto {
        if c == ' ' {
            if !espaco {
                fmt.Print(" ")
                espaco = true
            }
        } else {
            fmt.Printf("%c", c)
            espaco = false
        }
    }
    fmt.Println()
}