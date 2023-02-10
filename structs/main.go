package main

import (
	"fmt"
	"strings"
)

// Struct é um tipo usado para estruturar dados complexos
// Além de permitir atributos, ela também pode ter funções acopladas
// A regra dos valores default vale para structs também, por isso existe
// uma reclamação grande que em go você precisa defender de "zero-valued"
// Exemplo: caso você tenha uma API que recebe um JSON e um campo booleano não é enviado
// ele chegará na sua struct como false.
type band struct {
	name       string
	firstAlbum string
	members    []string
}

func (b *band) Repr() {
	fmt.Printf("Band: %s\nFirst Album: %s\nMembers: %s\n", b.name, b.firstAlbum, b.Members())
}

func (b *band) Members() string {
	return strings.Join(b.members, ", ")
}

func main() {
	blink182 := band{
		name:       "Blink-182",
		firstAlbum: "Cheshire Cat",
		members:    []string{"Mark Hoppus", "Travis Barker", "Tom DeLonge"},
	}
	blink182.Repr()
}
