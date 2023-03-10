package main

import "fmt"

// Coleções são maneiras de agrupar dados
// Em go temos dois tipos para lidar com coleções: arrays e mapas
// E os slices? Aguarde e verás...

func main() {
	// Array é um tipo que é inicializado com um total de valores que ele terá
	var names [3]string
	names[0] = "Mauricio"
	names[1] = "Lais"
	names[2] = "Bentinho"
	fmt.Printf("My name: %s\n", names[0])
	fmt.Println("All names:", names)
	// Arrays também podem ser inicializados "inline"
	otherNames := [3]string{"Mickey", "Minie", "Pluto"}
	fmt.Println(otherNames)
	// Podemos omitir o tamanho do array
	plusOtherNames := [...]string{"João", "Maria"}
	fmt.Println(plusOtherNames)

	// Slices são, em essência, arrays. Porém, são arrays que podem crescer
	// Um slice aponta pra arrays - é um "wrap" sobre arrays
	// Arrays não podem ter um total de elementos maior do que o definido em sua
	// inicialização
	// Esse conceito também existe em outras linguagens, como Rust com arrays e vetores
	// Para definir um slice, podemos omitir seu tipo
	colors := []string{"Blue", "Brown", "Green", "Red"}
	fmt.Println("Cores:", colors)

	// É possível iterar em arrays e slices usando for range
	for idx, color := range colors {
		fmt.Printf("Color at %d index: %s\n", idx, color)
	}

	// Podemos fatiar um slice, usando a seguinte sintaxe
	firstTwoColors := colors[0:2]
	fmt.Println("First two colors:", firstTwoColors)
	lastTwoColors := colors[2:4]
	fmt.Println("Last two colors:", lastTwoColors)

	// É possível criar um slice também usando make
	// make é uma função do go usada para criar alguns tipos como slices, maps e channels
	primeNumbers := make([]int, 0)
	primeNumbers = append(primeNumbers, 2, 3, 5, 7, 11, 13, 17, 19)
	fmt.Println("Prime numbers:", primeNumbers)

	// Mapas são como dicionários do Python ou Hash do Ruby
	// Esse map é um dicionário com chaves string e valores inteiros
	// Em go, os dicionários não podem ter tipos misturados
	// Importante: mapas em go não são ordenados
	users := map[string]int{
		"Mauricio": 32,
		"Lais":     29,
		"Bentinho": 2,
	}
	fmt.Printf("My name is %s and my age is %d\n", "Mauricio", users["Mauricio"])
	// Podemos iterar no par chave/valor como em outras linguagens
	for name, age := range users {
		fmt.Printf("Name %s and age %d\n", name, age)
	}
	// Podemos incluir mais uma usuário no mapa usando uma sintaxe similar a outras linguagens
	users["Catarina"] = 4
	// E imprimir de uma forma que ajuda a depuração do código
	fmt.Printf("Users: %#v\n", users)
	// Podemos deletar usando a função delete
	// Essa função deleta usando a referência
	delete(users, "Catarina")
	// Testando se ainda existe o elemento
	// A idade retorna 0 porque é o valor default do inteiro
	// Diferente de Rust, Go não tem recursos para mostrar a ausência de um valor
	catarinaAge, ok := users["Catarina"]
	fmt.Printf("Usuário existe? %t Se sim, qual é a idade? %d\n", ok, catarinaAge)
}
