package main
import "fmt"
func main() {
    var n1, n2, n3 int
    fmt.Scan(&n1)
    fmt.Scan(&n2)
    fmt.Scan(&n3)

    if(n1 == n2 && n1 == n3){
        fmt.Printf("3\n")
    } else if(n2 == n1 || n2 == n3 || n1 == n3){
        fmt.Printf("2\n")
    } else{
        fmt.Printf("0\n")
    }

}
