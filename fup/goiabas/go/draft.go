package main
import "fmt"
func main() {
    var C int
    fmt.Scan(&C)
    var qtd_bananas int
    var qtd_goiabas int
    var qtd_mangas int
    fmt.Scan(&qtd_bananas, &qtd_goiabas, &qtd_mangas)

    total_frutas := qtd_bananas + qtd_goiabas + qtd_mangas

    if total_frutas % C == 0 {
        viagems := total_frutas / C
        fmt.Println(viagems)
    } else {
        viagems := (total_frutas / C) + 1
        fmt.Println(viagems)
    }
}