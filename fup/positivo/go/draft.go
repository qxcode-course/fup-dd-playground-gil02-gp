package main
import "fmt"
func main() {
    var a int
    fmt.Scan(&a)
    if(a >= 0){
        fmt.Printf("SIM\n")
    } else{
        fmt.Printf("NAO\n")
    }
}
