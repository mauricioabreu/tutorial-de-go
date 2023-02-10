package pong

import "fmt"

// Função que não pode ser chamada fora do pacote
func hello() {
	fmt.Println("Pong")
}

func Hello() {
	fmt.Println("Pong")
}
