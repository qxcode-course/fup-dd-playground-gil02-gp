package main
import "fmt"
func main() {
    var a, b int
    fmt.Scan(&a, &b)
    
    diferente := (b - a + 15) % 15
    if diferente == 0 {
        fmt.Println("Empate")
    } else if diferente <= 7 {
        fmt.Println("Jogador 1")
    } else {
        fmt.Println("Jogador 2")
    }
}
