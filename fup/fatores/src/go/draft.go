package main
import "fmt"
func main() {
    var n int
    fmt.Scan(&n)
    fator := 2
    cont := 0

    for n != 1{
        if n % fator == 0{
            n /= fator
            cont++
        } else {
            if cont > 0{
                fmt.Printf("%d %d\n", fator, cont)
            }
            fator++
            cont = 0
        }
    }
    if cont > 0{
        fmt.Printf("%d %d\n", fator, cont)
    }
}
