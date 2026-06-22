package main
import (
    "fmt"
    "bufio"
    "os"
    "strings"
)
func main() {
    entrada := bufio.NewReader(os.Stdin)
    var t int
    fmt.Fscan(entrada, &t)

    for ; t > 0; t-- {
        var ultron, pessoa string
        fmt.Fscan(entrada, &ultron, &pessoa)

        ultron = strings.TrimSpace(ultron)
        pessoa = strings.TrimSpace(pessoa)

        var existir[26] bool
        for _, c := range ultron {
            existir[c-'a'] = true
        }
        iguais := 0
        for _, c := range pessoa {
            if existir[c-'a'] {
                iguais++
            }
        }
        if iguais == len(pessoa) {
            fmt.Println("chefe")
        } else if iguais*2 > len(pessoa) {
            fmt.Println("ultron")
        } else {
            fmt.Println("pessoa")
        }
    }
}