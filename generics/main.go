package main

import "fmt"

// Podemos criar uma interface pra ter um tipo bem definido como uma constraint
type Number interface {
	int64 | float64 | float32
}

// Generics é uma forma de você ter uma mesma API para tipos de dados diferentes, especialmente
// tipos concretos
// Por exemplo, se quisermos somar uma lista de valores, anteriormente em go você tinha que ter
// uma função para cada tipo. Somar uma lista de floats e somar uma lista de inteiros dependia
// de ter funções iguais com assinaturas diferentes
func nonGenericSumInt(numbers []int) int {
	total := 0
	for _, number := range numbers {
		total += number
	}

	return total
}

func nonGenericSumFloat(numbers []float64) float64 {
	total := 0.0 // Aqui precisa ser 0.0 porque o go assume 1 como inteiro nesse contexto
	for _, number := range numbers {
		total += number
	}

	return total
}

// Essa é uma função que lida com tipo parametrizável. As maiores diferenças entre código com
// generics e interfaces são essas três:
// 1) Generics são mais performáticas e seguras pois não precisam de type assertion
// 2) Todo código é otimizado em tempo de compilação
// 3) Interface é mais sobre contrato e implementação do que código genérico
func genericSum[T int | float64 | float32](numbers []T) T {
	var total T
	for _, value := range numbers {
		total += value
	}
	return total
}

func main() {
	fmt.Println("Non generic sum:", nonGenericSumInt([]int{1, 2, 3}))
	fmt.Println("Non generic sum:", nonGenericSumFloat([]float64{1, 2, 3}))
	fmt.Println("Generic sum of integers:", genericSum([]int{1, 2, 3}))
	fmt.Println("Generic sum of floats:", genericSum([]float64{1, 2, 3}))
}
