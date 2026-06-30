package main
import "fmt"
func main() {
    cartela := [4][4] int {
        {1, 9, 27, 23},
        {34, 20, 37, 47},
        {30, 87, 55, 69},
        {13, 60, 99, 66},
    }
    acerto := 0
    for i := 0; i < 6; i++ {
        var num int
        fmt.Scan(&num)

        encontrou := false

        for j := 0; j < 4 && !encontrou; j++ {
            for c := 0; c < 4; c++ {
                if cartela[j][c] == num {
                    acerto++
                    encontrou = true
                    break
                }
            }
        }
    }
    fmt.Println(acerto)
}