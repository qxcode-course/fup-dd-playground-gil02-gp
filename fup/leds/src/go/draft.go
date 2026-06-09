package main
import "fmt"
func main() {
    var painel[101] string
    var n int
    var qtd = 0
    fmt.Scan(&n)

    for j := 0; j < n; j++ {
        qtd = 0
        fmt.Scan(&painel[j])
    }
    for i := 0; i < n; i++ {
    if (painel[i] == "0"){
        qtd += 6
    }
    if (painel[i] == "1"){
        qtd += 2
    } 
    if (painel[i] == "2"){
        qtd += 5
    } 
    if (painel[i] == "3"){
        qtd += 5
    } 
    if (painel[i] == "4"){
        qtd += 4
    } 
    if (painel[i] == "5"){
        qtd += 5
    }
    if (painel[i] == "6"){
        qtd += 6
    }
    if (painel[i] == "7"){
        qtd += 3
    }
    if (painel[i] == "8"){
        qtd += 7
    }
    if (painel[i] == "9"){
        qtd += 6
    }
    }
    fmt.Printf("%d leds\n", qtd)
}