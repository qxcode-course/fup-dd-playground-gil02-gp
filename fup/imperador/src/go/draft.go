package main

import "fmt"

func main() {
	var n int
	fmt.Scan(&n)

	mat := make([][]byte, n)

	linhaLeao := -1
	colunaLeao := -1
	for i := 0; i < n; i++ {
		mat[i] = make([]byte, n)
		for j := 0; j < n; j++ {
			fmt.Scanf(" %c", &mat[i][j])
			if mat[i][j] == 'L' {
				linhaLeao = i
				colunaLeao = j
				continue
			}
		}
	}
	gladiador := 0
	condenado := 0

	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			if linhaLeao != -1 && (i == linhaLeao || j == colunaLeao) {
				continue
			}
			if mat[i][j] == 'G' {
				gladiador += 2
			}
			if mat[i][j] == 'C' {
				condenado++
			}
		}
	}
	if gladiador > condenado {
		fmt.Println("Gladiadores")
	} else if condenado > gladiador {
		fmt.Println("Condenados a morte")
	} else {
		fmt.Println("Ninguem")
	}
}
