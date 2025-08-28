package internal

import "fmt"

func Calculadora(a, b float64) {
	fmt.Println("Calculando...")
	Somar := a + b
	fmt.Printf("A Soma de %.0f e %.0f é: %.0f\n", a, b, Somar)
	Subtracao := a - b
	fmt.Printf("A Subtração de %.0f e %.0f é: %.0f\n", a, b, Subtracao)
	Multiplicacao := a * b
	fmt.Printf("A Multiplicação de %.0f e %.0f é: %.0f\n", a, b, Multiplicacao)

	if b == 0 {
		fmt.Println("Erro: Você está tentando dividir por zero")
	} else {
		Divisao := a / b
		if Divisao >= 0 {
			fmt.Printf("A divisão teve um resultado posivito %.0f\n", Divisao)
		} else {
			fmt.Printf("A divisão teve um resultado negativo %.0f\n", Divisao)
		}
		Resto := int(a) % int(b)
		fmt.Printf("O Resto de %.0f e %.0f é: %d\n", a, b, Resto)
	}
}

func CalcAreaRetangulo(base, altura float64) {
	Area := base * altura
	fmt.Printf("A Area do retângulo é: %.2f\n", Area)
}

