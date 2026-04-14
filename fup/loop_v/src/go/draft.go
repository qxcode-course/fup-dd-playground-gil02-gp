package main
import "fmt"
func main() {
    var a, b int
    fmt.Scan(&a, &b)
    
    i := a
    primeiro := true
    fmt.Printf("[ ")
    for {
        if i >= b {
            break
        }

        if i%2 == 0 {
             i++
             continue
        }

        if ! primeiro {
            fmt.Print(" ")
        }
        fmt.Print(i)
        primeiro = false
        i++
    }
    fmt.Printf(" ]\n")
}
