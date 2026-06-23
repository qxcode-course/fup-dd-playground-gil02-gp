package main
import (
    "fmt"
    "bufio"
    "os"
    "strings"
)
func main() {
    entrada := bufio.NewReader(os.Stdin)
    var nIron, nCaptain int
    fmt.Fscan(entrada, &nIron)
    entrada.ReadByte()

    somaIron := 0
    somaCaptain := 0

    campeao := ""
    maior := -1

    for i := 0; i < nIron; i++ {
        nome, _ := entrada.ReadString('\n')
        nome = strings.TrimSpace(nome)

        var poder int
        fmt.Fscan(entrada, &poder)
        entrada.ReadByte()
        somaIron += poder
        if poder > maior {
            maior = poder
            campeao = nome
        }
    }
    fmt.Fscan(entrada, &nCaptain)
    entrada.ReadByte()

    for i := 0; i < nCaptain; i++ {
        nome, _ := entrada.ReadString('\n')
        nome = strings.TrimSpace(nome)
        var poder int
        fmt.Fscan(entrada, &poder)
        entrada.ReadByte()

        somaCaptain += poder
        if poder > maior {
            maior = poder
            campeao = nome
        }
    }
    if somaCaptain > somaIron {
        fmt.Println("Team Captain Wins")
    } else if somaIron > somaCaptain {
        fmt.Println("Team Iron Wins")
    } else {
        fmt.Println("Draw")
    }
    fmt.Println(campeao)
}