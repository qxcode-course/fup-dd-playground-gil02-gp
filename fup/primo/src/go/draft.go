package main
import "fmt"
func primo(n int) int {
    if n < 2 {
        return 0
    }
    for i := 2; i <= n/2; i++ {
        if n%i == 0 {
            return 0
        }
    }
    return 1
}
func main() {
    var n int
    fmt.Scan(&n)
    fmt.Println(primo(n))
}