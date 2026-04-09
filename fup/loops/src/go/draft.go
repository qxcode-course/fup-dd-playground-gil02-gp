package main
import "fmt"
func main() {
    var angulo int
    fmt.Scan(&angulo)

    angulo = angulo % 360
    if angulo < 0 {
        angulo += 360
    }
    fmt.Println(angulo)
}
