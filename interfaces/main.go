package main

import "fmt"

// Interface é uma maneira de declararmos um conjunto de comportamentos
// que uma struct deve obedecer. Em go, elas são extremamente úteis por dois motivos:
// Histórico, porque go até ontem não tinha generics
// Porque interfaces são maneiras boas de definirmos comportamentos que um determinado tipo
// deve implementar, o que facilita passar diferentes implementações para uma mesma função

// A definição da interface precisa ter o nome das funções
type animal interface {
	makeNoise() string
	// Descomentar a função abaixo gera um erro pois as structs não implementam
	// a função totalPaws() int
	// totalPaws() int
}

// Agora criamos structs que implementam essa interface
type dog struct{}

func (d *dog) makeNoise() string {
	return "Woof!"
}

type cat struct{}

func (c *cat) makeNoise() string {
	return "Meow!"
}

func goAnimalMakeANoise(pet animal) {
	fmt.Printf("What is your noise? %s\n", pet.makeNoise())
}

func main() {
	// Inicializamos as structs e chamamos a função goAnimalMakeANoise
	// com structs diferentes.
	d := &dog{}
	goAnimalMakeANoise(d)

	c := &cat{}
	goAnimalMakeANoise(c)
}
