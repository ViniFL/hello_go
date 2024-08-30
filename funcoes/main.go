package main

import (
	"fmt"
)

type Carro struct {
	Nome string
}

func (c Carro) andar() {
	fmt.Println(c.Nome, "Andou")
}

func main() {

	carro := Carro{
		Nome: "Hilux",
	}

	carro.andar()

	resultado := func(x ...int) func() int {
		res := 0
		for _, v := range x {
			res += v
		}
		return func() int {
			return res * res
		}
	}

	// resultado := somaTudo(10, 20, 257, 378)

	fmt.Println(resultado(45, 567, 213, 643)())
}

func soma(a int, b int) (result int) {
	result = a + b
	return
}

func somaTudo(x ...int) int {
	resultado := 0

	for _, v := range x {
		resultado += v
	}
	return resultado
}
