package main
import "fmt"
func ehVogal(c byte) bool {
    return c == 'a' || c == 'e' || c == 'i' || c == 'o' || c == 'u' ||
    c == 'A' || c == 'E' || c == 'I' || c == 'O' || c == 'U'
}
func contaVogal(s string) int {
    cont := 0
    for i := 0; i < len(s); i++ {
        if ehVogal(s[i]){
            cont++
        }
    }
    return cont
}
func main() {
    var n1, n2 string
    fmt.Scan(&n1, &n2)

    ponto := 0

    if n1[0] == n2[0] {
        ponto += 20
    }
    if len(n1) == len(n2) {
        ponto += 30
    }
    if contaVogal(n1) == contaVogal(n2) {
        ponto += 30
    }
    fim1 := ehVogal(n1[len(n1)-1])
    fim2 := ehVogal(n2[len(n2)-1])

    if fim1 == fim2 {
        ponto += 20
    } else {
        ponto -= 10
    }
    if ponto < 0 {
        ponto = 0
    }
    fmt.Printf("As chances do crush te dar bola sao: %d%%!\n", ponto)
}