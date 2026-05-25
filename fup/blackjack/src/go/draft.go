package main
import "fmt"
func main() {
    var tam int
    fmt.Scan(&tam)
    count := 0
    as := 0
    var x int
    for i := 0; i < tam; i++ {
        fmt.Scan(&x)
        if x == 1 {
            as++
            count +=11
        } else {
            if x >= 11 && x <= 13 {
                count += 10
            } else {
                count += x
            }
        }
    }
    for i := 0; i < as; i++ {
        if count > 21 {
            count -= 10
        }
    }
    fmt.Println(count)
}
