package main
import "fmt"
func main() {
    var nota1, nota2, nota_final float64
    fmt.Scan(&nota1, &nota2, &nota_final)
    media := (nota1 + nota2) / 2

    if(media >= 7){
        fmt.Printf("aprovado\n")
    } else if(media < 4){
        fmt.Printf("reprovado\n")
    }else if(nota_final >= 5){
        fmt.Printf("aprovado na final\n")
    } else {
        fmt.Printf("reprovado na final\n")
    }

}
