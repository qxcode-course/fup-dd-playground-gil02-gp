package main
import "fmt"
func main() {
    var n int
    var pe string
    fmt.Scan(&n)
    fmt.Scan(&pe)
    fmt.Print("[ ")
    for i := 0; i <= 10; i++ {
        if i == n {
            continue
        }
        if i == 10{
            fmt.Print("ceu ")
        } else {
            fmt.Printf("%d%s ", i, pe)
        }
        if pe == "e"{
            pe = "d"
        } else{
            pe = "e"
        }
    }
    fmt.Print("]\n")
}
