package main
import "fmt"
func main() {
    var B, T int
    fmt.Scan(&B, &T)

    felix := ((B + T) * 70) / 2
    marzia := (160 * 70) - felix

    if felix > marzia {
        fmt.Println("1")
    } else if marzia > felix {
        fmt.Println("2")
    } else{
        fmt.Println("0")
    }
}
