package main

import (
	"github.com/mauricioabreu/tutorial-de-go/packages/ping"
	"github.com/mauricioabreu/tutorial-de-go/packages/pong"
)

// Em go, para usar funções de um pacote, elas precisam estar exportadas.
// Para exportar um recurso, seja variável, struct, função, etc, precisam estar com a primeira letra
// em maiúsculo.
// Essa é uma maneira de manter baixo acomplamento entre pacotes e trazer uma
// maior fluidez na comunicação.
func main() {
	ping.Hey()
	// A função abaixo existe, mas está com a primeira letra minúscula
	// pong.hello()
	pong.Hello()
}
