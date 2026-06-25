package main
import "fmt"
type Jogada struct {
    p1, p2 int
}
func calc_pontuacao (j Jogada) (bool, int) {
    if j.p1 < 10 || j.p2 < 10 {
        return false, 0
    }
    ponto := j.p1 - j.p2
    if ponto < 0 {
        ponto = -ponto
    }
    return true, ponto
}
    func procurar_melhor_jogada(jogadas []Jogada) int {
        melhorIndice := -1
        MenorPontuacao := 0

        for i := 0; i < len(jogadas); i++{
            valida, pontuacao := calc_pontuacao(jogadas[i])
            if !valida {
                continue
            }
            if melhorIndice == -1 || pontuacao < MenorPontuacao {
                melhorIndice = i
                MenorPontuacao = pontuacao
            }
        }
        return melhorIndice
    }

func main() {
    var n int
    fmt.Scan(&n)

    jogadas := make([]Jogada, n)
    for i := 0; i < n; i++ {
        fmt.Scan(&jogadas[i].p1, &jogadas[i].p2)
    }
    indice := procurar_melhor_jogada(jogadas)
    if indice == -1 {
        fmt.Println("sem ganhador")
    }else {
        fmt.Println(indice)
    }
}