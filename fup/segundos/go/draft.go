package main
import "fmt"
func main() {
    var tempo int
    fmt.Scan(&tempo)

    var resto int = tempo % 3600
    var hora = tempo / 3600
    var minutos = resto / 60
    var segundo = resto % 60

    fmt.Printf("%d:%d:%d", hora, minutos, segundo)
}
