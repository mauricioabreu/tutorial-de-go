package main

import (
	"errors"
	"fmt"
)

// Um erro é uma maneira de comunicar que algo na função deu errado
// Em go, existe um pacote errors destinado a ajudar a cuidar de erros

// Erros podem ser comparáveis também, para tomar a decisão correta quando
// um erro acontecer. Por exemplo, é importante que sua lib tenha os erros exportados
// para que clientes da lib possam saber o que fazer com cada tipo de erro
var (
	ErrIDontLikeNumber2 = errors.New("number 2 not allowed")
	ErrIDontLikeNumber3 = errors.New("number 3 not allowed")
)

func doSomething(status bool) error {
	if status {
		return nil
	}
	return errors.New("bad status")
}

func withDifferentErrors(number int) error {
	if number == 2 {
		return ErrIDontLikeNumber2
	}
	if number == 3 {
		return ErrIDontLikeNumber3
	}
	return nil
}

func main() {
	// Esse é um jeito de verificar se a função tem erro
	// Mas poderia ser também
	// err := doSomething(true)
	// if err != nil {
	//     fmt.Println("Sem erro")
	// }
	if err := doSomething(true); err != nil {
		fmt.Println("Tem erro?", err)
	}
	if err := doSomething(false); err != nil {
		fmt.Println("Tem erro?", err)
	}
	err := withDifferentErrors(2)
	if err == ErrIDontLikeNumber2 {
		fmt.Println("Number not allowed:", err)
	}
	if err == ErrIDontLikeNumber3 {
		fmt.Println("Number not allowed:", err)
	}

	// É possível também verificar se um erro faz parte de outro
	internalError := errors.New("bad input for function")
	anotherError := errors.New("any error")
	wrappedError := fmt.Errorf("something went wrong: %w", internalError)
	if errors.Is(wrappedError, internalError) {
		fmt.Println("Yes, error is wrapped with internalError")
	}
	if errors.Is(wrappedError, anotherError) {
		fmt.Println("Yes, error is wrapped with anotherError")
	}
}
