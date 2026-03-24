package main
import "fmt"
func main() {
    var num1, num2 int
    var operador string
    fmt.Scan(&num1)
    fmt.Scan(&num2)
    fmt.Scan(&operador)

    if(operador == "+"){
        fmt.Println(num1 + num2)
    } else if(operador == "-"){
        fmt.Println(num1 - num2)
    } else if(operador == "*"){
        fmt.Println(num1 * num2)
    } else if(operador == "/"){
        fmt.Println(num1 / num2)
    }
}
