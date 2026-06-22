package main
import "fmt"
func main() {
    var letra string
    var rot int
    fmt.Scan(&letra)
    fmt.Scan(&rot)

    pos := int(letra[0] - 'a')
    pos = (pos + rot) % 26

    if pos < 0 {
        pos += 26
    }
    fmt.Printf("%c\n", byte(pos)+'a')
}