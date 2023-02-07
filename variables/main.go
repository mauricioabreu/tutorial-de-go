package main

import "fmt"

func main() {
	// Variáveis em go podem ser definidas de dois jeitos
	// Usando var
	var me string
	me = "Mauricio"
	age := 32

	// Imprimir valores com interpolação com uso de funções que terminam em 'f'
	// Exemplos: Sprintf, Printf, Msgf, etc
	fmt.Printf("Name: %s and I'm %d years old\n", me, age)

	// Variáveis em go tem tipos
	pi := 3.14159265359
	fmt.Printf("Pi number: %f\n", pi)
	active := true
	fmt.Printf("Active? %t\n", active)

	// Coleções tem tipos
	var people []string
	people = append(people, me)
	girlfriend := "Lais"
	people = append(people, girlfriend)
	fmt.Printf("People: %v\n", people)
}
