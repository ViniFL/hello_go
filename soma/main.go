package main

import (
	"fmt"
	"soma/math"
)

func main() {
	resultado := math.Soma(29, 78)
	fmt.Printf("%v", resultado)

	resultadoX := math.SomaX(59)
	fmt.Printf("%v", resultadoX)
}
