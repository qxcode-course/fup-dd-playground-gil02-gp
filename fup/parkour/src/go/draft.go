package main
import "fmt"
func main() {
    var n int
    fmt.Scan(&n)

    var bloco[50] int
    for i := 0; i < n; i++ {
        fmt.Scan(&bloco[i])
    }
    parkour := 0
    for i := 1; i < n; i++ {
        if (bloco[i]-bloco[i-1]) > 1 || (bloco[i]-bloco[i-1]) < -1 {
            parkour ++
        }
    }
    fmt.Println(parkour)
}
