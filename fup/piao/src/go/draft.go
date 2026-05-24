package main

import "fmt"

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}
func main() {
	var l, n int
	fmt.Scan(&l)
	fmt.Scan(&n)

	ganhador := -1
	melhor := l + 1
	perdedor := 0
	pior := -1

	for i := 0; i < n; i++ {
		var jogada int
		fmt.Scan(&jogada)
		dista := abs(jogada)

		if dista <= l && dista <= melhor {
                melhor = dista
                ganhador = i
		}
        
		if dista >= pior {
                pior = dista
                perdedor = i
        }
	}

	if ganhador == -1 {
		fmt.Println("nenhum")
	} else {
		fmt.Println(ganhador)
	}

	fmt.Println(perdedor)
}
