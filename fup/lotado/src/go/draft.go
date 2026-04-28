package main
import "fmt"
func main() {
    var c, m int
    fmt.Scan(&c)
    passageiro := 0

    for passageiro < 2 * c {
        fmt.Scan(&m)
        passageiro += m
        if passageiro >= 2 * c{
        fmt.Printf("hora de partir\n")
        return
    }else if passageiro == 0 {
        fmt.Printf("vazio\n")
    } else if passageiro < c {
        fmt.Printf("ainda cabe\n")
    } else {
        fmt.Printf("lotado\n")
    }
        
    }
    
}
