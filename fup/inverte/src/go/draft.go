package main
import "fmt"
func main() {
    var c byte
    fmt.Scanf("%c", &c)

    if c >= 'a' && c <= 'z' {
        fmt.Printf("%c\n", c - 'a' + 'A')
    } else if c >= 'A' && c <= 'Z' {
        fmt.Printf("%c\n", c - 'A' + 'a')
    } else {
        fmt.Printf("%c\n", c)
    }
}