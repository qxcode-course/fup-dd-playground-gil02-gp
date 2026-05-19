package main

import "fmt"

func main() {
	var n int
	fmt.Scan(&n)
	vetor := make([]int, n)
	for i := 0; i < n; i++ {
		fmt.Scan(&vetor[i])
	}
	fmt.Print("[")
	for i, v := range vetor {
		var carta string
		switch v {
		case 1:
			carta = "A"
		case 11:
			carta = "J"
		case 12:
			carta = "Q"
		case 13:
			carta = "K"
		default:
			carta = fmt.Sprintf("%d", v)
		}
		if i > 0 {
			fmt.Print(", ")
		}
		fmt.Print(carta)
	}
	fmt.Println("]")

}
