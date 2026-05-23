package main
import "fmt"
func main() {
    var n int
    fmt.Scan(&n)

    var esquerda[61] int
    var direita[61] int

    for i := 0; i < n; i++ {
        tamanho := 0
        var lado string
        fmt.Scan(&tamanho, &lado)

        if lado == "E" {
            esquerda[tamanho]++
        } else if lado == "D" {
            direita[tamanho]++
        }
    }
    total := 0
    for i := 0; i < 61; i++ {
        if esquerda[i] > direita[i] {
            total += direita[i]
        } else {
            total += esquerda[i]
        }
    }
    fmt.Println(total)
}
