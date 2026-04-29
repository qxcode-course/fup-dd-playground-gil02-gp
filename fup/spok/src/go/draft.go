package main
import "fmt"
func main() {
    var id, n int
    fmt.Scan(&id)

    if id < 0{
        fmt.Println("0")
    }

    n = id
    reverso := 0

    for n > 0{
        reverso = reverso * 10 + (n % 10)
        n /= 10
    }
    if id == reverso{
        fmt.Println("1")
    } else {
        fmt.Println("0")
    }
}
