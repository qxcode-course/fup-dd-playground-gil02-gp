package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	entrada := bufio.NewReader(os.Stdin)
	texto, _ := entrada.ReadString('\n')
	if len(texto) > 0 && texto[len(texto)-1] == '\n' {
		texto = texto[:len(texto)-1]
	}
	op, _ := entrada.ReadByte()
	switch op {
	case 'M':
		for i := 0; i < len(texto); i++ {
			c := texto[i]
			if c >= 'a' && c <= 'z' {
				c -= 'a' - 'A'
			}
			fmt.Printf("%c", c)
		}
	case 'm':
		for i := 0; i < len(texto); i++ {
			c := texto[i]
			if c >= 'A' && c <= 'Z' {
				c += 'a' - 'A'
			}
			fmt.Printf("%c", c)
		}
	case 'i':
		for i := 0; i < len(texto); i++ {
			c := texto[i]
			if c >= 'a' && c <= 'z' {
				c -= 'a' - 'A'
			} else if c >= 'A' && c <= 'Z' {
				c += 'a' - 'A'
			}
			fmt.Printf("%c", c)
		}
	case 'p':
		inicio := true
		for i := 0; i < len(texto); i++ {
			c := texto[i]
			if c == ' ' {
				inicio = true
				fmt.Printf("%c", c)
				continue
			}
			if inicio {
				isVogal := c == 'a' || c == 'e' || c == 'i' || c == 'o' || c == 'u' ||
					c == 'A' || c == 'E' || c == 'I' || c == 'O' || c == 'U'
				isUltimaLetra := (i+1 == len(texto) || texto[i+1] == ' ')

				if isVogal && isUltimaLetra {
					if c >= 'A' && c <= 'Z' {
						c += 'a' - 'A'
					}
				} else if c >= 'a' && c <= 'z' {
					c -= 'a' - 'A'
				} else if c >= 'A' && c <= 'Z' {
				}
				inicio = false
			} else {
				if c >= 'A' && c <= 'Z' {
					c += 'a' - 'A'
				}
			}
			fmt.Printf("%c", c)
		}
	}
	fmt.Println()
}
