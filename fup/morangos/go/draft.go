package main
import "fmt"
func main() {
    var  c1, l1, c2, l2 float64
    fmt.Scan(&c1, &l1, &c2, &l2)
    area1 := c1 * l1
    area2 := c2 * l2

    if area1 > area2 {
        fmt.Println(area1)
    } else {
        fmt.Println(area2)
    } 
}
