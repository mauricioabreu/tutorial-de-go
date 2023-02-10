package main

import "fmt"

// Funções são definidas com nome e parâmetros
// Elas também poder ter retornos
// Exemplo:
//
//	func helloWorld(name string) string {
//		return "Hello, " + name
//	}
func helloWorld(name string) {
	fmt.Printf("Hello, %s\n", name)
}

// Funções podem aceitar funções como parâmetro
// HOF - higher order functions
// As funções precisam ter exatamente a mesma assinatura da função
// que você deseja passar ao ser chamada
// É possível colocar uma função até como campo de uma struct
func welcome(name string, helloWorldF func(string)) {
	helloWorldF(name)
}

// Funções podem receber uma sequência de argumentos do mesmo tipo
// Por exemplo, funções
func sumNumbers(numbers ...int) int {
	total := 0
	for _, number := range numbers {
		total += number
	}
	return total
}

func main() {
	helloWorld("Mauricio")
	// Também é possível definir funções anônimas
	// \x -> x + 1 (Haskell)
	// lambda x : x + 1 (Python)
	helloWorldFunc := func(name string) {
		fmt.Printf("Hello, %s\n", name)
	}
	helloWorldFunc("Mauricio")
	welcome("Mauricio", helloWorldFunc)
	fmt.Println("Sum numbers:", sumNumbers(1, 2))
	fmt.Println("Sum numbers:", sumNumbers(10, 20, 30))
}
