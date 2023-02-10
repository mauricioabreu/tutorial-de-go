package main

import "fmt"

// Em go, temos duas maneiras de fazer controle de fluxo: if e switch
func main() {
	// if/else tem a seguinte sintaxe
	number := 10
	if number > 10 {
		fmt.Printf("Number %d is greater than 10", number)
	} else {
		fmt.Printf("Number %d not greater than 10", number)
	}

	// Switch pode ser avaliado em torno de uma variável
	switch number {
	case 9:
		fmt.Println("9")
	case 10:
		fmt.Println("10")
	}

	// Ou em torno de comparações mais livres
	// O switch do go pára na primeira avaliação em que o resultado é uma verdade
	active := true
	switch {
	case number > 5:
		fmt.Println("Number 5 is greater")
	case active:
		fmt.Println("true is true")
	}
}
