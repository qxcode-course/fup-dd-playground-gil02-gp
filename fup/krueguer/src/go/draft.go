package main
import (
    "fmt"
    "bufio"
    "os"
)
func vogal(c byte) bool {
    return c == 'a' || c == 'e' || c == 'i' || c == 'o' || c == 'u'
}
func main() {
    entrada := bufio.NewReader(os.Stdin)
    var n int
    fmt.Fscan(entrada, &n)
    entrada.ReadBytes('\n')

    for ; n > 0; n-- {
        var s string
        s, _ = entrada.ReadString('\n')

        if len(s) > 0 && s[len(s)-1] == '\n' {
            s = s[:len(s)-1]
        }
        inicioAtual := 0
        tamAtual := 0
        inicioMaior := 0
        tamMaior := 0

        for i := 0; i < len(s); i++ {
            if vogal(s[i]) {
                if tamAtual == 0 {
                    inicioAtual = i
                }
                tamAtual++
            } else {
                if tamAtual > tamMaior {
                    inicioMaior = inicioAtual
                    tamMaior = tamAtual
                }
                tamAtual = 0
            }
        }
        if tamAtual > tamMaior {
            inicioMaior = inicioAtual
            tamMaior = tamAtual
        }
        fmt.Printf(s[inicioMaior:inicioMaior+tamMaior] + "\n")
    }
    
}