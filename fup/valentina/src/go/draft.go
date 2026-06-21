package main

import "fmt"

func main() {
	var s1, s2, opStr string
	if _, err := fmt.Scan(&s1, &opStr, &s2); err != nil {
		fmt.Println("Erro ao ler a entrada:", err)
		return
	}
	if len(s1) == 0 || len(s2) == 0 || len(opStr) == 0 {
		fmt.Println("Entrada inválida")
		return
	}

	n1 := s1[0]
	n2 := s2[0]
	op := opStr[0]

	a := int(n1 - 'a')
	b := int(n2 - 'a')

	var resultado int
	if op == '+' {
		resultado = (a + b) % 26
	} else {
		resultado = (a - b + 26) % 26
	}
	fmt.Printf("%c\n", byte(resultado)+'a')
}
