package main
import "fmt"
func main() {
    var n1, n2, n3, n4 int
    fmt.Scan(&n1)
    fmt.Scan(&n2)
    fmt.Scan(&n3)
    fmt.Scan(&n4)

    if n1 >= n2 && n1 >= n3 && n1 >= n4 {
        fmt.Println(n1)
    } else if n2 >= n1 && n2 >= n3 && n2 >= n4 {
        fmt.Println(n2)
    } else if n3 >= n1 && n3 >= n2 && n3 >= n4 {
        fmt.Println(n3)
    } else {
        fmt.Println(n4)
    }
}
