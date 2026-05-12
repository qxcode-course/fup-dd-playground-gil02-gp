package main
import "fmt"
func main() {
    var n1, n2 int
    fmt.Scan(&n1)
    var vetor [50]int
    for i := 0; i < n1; i++{
        fmt.Scan(&vetor[i])
    }
    fmt.Scan(&n2)
    var vetor2 [50]int
    for i := 0; i < n2; i++{
        fmt.Scan(&vetor2[i])
    }
    contido := true

    for i := 0; i < n1; i++{
        num := false
        for j := 0; j < n2; j++{
            if vetor[i] == vetor2[j] {
                num = true
                break
            }
        }
        if !num {
            contido = false
            break
        }
    }
    if contido {
        fmt.Printf("sim\n")
    } else {
        fmt.Printf("nao\n")
    }
}
