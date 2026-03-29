package main
import "fmt"
func main() {
    var n1, n2, n3 int
    fmt.Scan(&n1)
    fmt.Scan(&n2)
    fmt.Scan(&n3)

    if(n1 > n2 && n1 < n3) || (n1 < n2 && n1 > n3){
        fmt.Println(n1)
    } else if (n2 > n1 && n2 < n3) || (n2 < n1 && n2 > n3){
        fmt.Println(n2)
    } else {
        fmt.Println(n3)
    }
}