package main
import "fmt"
func main() {
    var a, b int
    fmt.Scan(&a) 
    fmt.Scan(&b)

    cont := 0
    if a == 0 && b == 0 {
        fmt.Println(1)
        return
    }
    for b > 0 {
        if b % 10 == a {
            cont++
        }
        b /= 10
    }

    fmt.Println(cont)
}