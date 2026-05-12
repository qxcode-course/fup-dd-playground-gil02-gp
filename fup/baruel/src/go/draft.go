package main
import "fmt"
func main() {
    var n, total int
    fmt.Scan(&total)
    fmt.Scan(&n)

    var figurinha[100] int
    for i := 0; i < n; i++ {
        fmt.Scan(&figurinha[i])
    }

    fmt.Print("[ ")
    for i := 0; i < n-1; i++ {
        if figurinha[i] == figurinha[i+1] {
            fmt.Print(figurinha[i], " ")
        }
    }
    fmt.Print("]\n")

    fmt.Print("[ ")
    for num := 1; num <= total; num++ {
        figu := false
        for i := 0; i < n; i++ {
            if num == figurinha[i] {
                figu = true
                break
            }
        }
        if !figu {
            fmt.Print(num, " ")
        }
    }
    fmt.Print("]\n")
}
