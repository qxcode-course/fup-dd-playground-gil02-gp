package main
import "fmt"
func main() {
    var numero int
    fmt.Scan(&numero)

    if(numero % 7 == 0){
        fmt.Printf("SIM\n")
    } else {
        fmt.Printf("NAO\n")
    }
}
