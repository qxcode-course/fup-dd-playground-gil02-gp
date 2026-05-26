package main

import "fmt"

func main() {
	var n int
	fmt.Scan(&n)

	melhor := -1
	melhorpontuacao := 101
	for i := 0; i < n; i++ {
		var a, b int
		fmt.Scan(&a, &b)
		if a >= 10 && b >= 10 {
			pontuacao := a - b
			if pontuacao < 0 {
				pontuacao = -pontuacao
			}
			if pontuacao < melhorpontuacao {
				melhorpontuacao = pontuacao
				melhor = i
			}
		}
	}
	if melhor == -1 {
		fmt.Println("sem ganhador")
	} else {
		fmt.Printf("%d\n", melhor)
	}
}
