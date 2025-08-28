package main

import (
	"fmt"
	"v2/internal"
)

func main() {
	Isabelle := "Meu amor"
	fmt.Println(internal.Nome)
	fmt.Println(internal.Idade)
	fmt.Println(internal.Altura)
	fmt.Println(internal.Estudante)
	fmt.Println(internal.Cidade)
	fmt.Println(Isabelle)

	internal.Calculadora(15,0)
	internal.CalcAreaRetangulo(15.5,4.2)
}