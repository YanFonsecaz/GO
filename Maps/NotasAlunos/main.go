package main

import "fmt"

func main() {
	NotasAlunos := map[string]float64{
		"Yan":    9.8,
		"Vini":   8.5,
		"Joao":   7.5,
		"Teste":  6.5,
		"Teste2": 5.5,
		"Teste3": 4.5,
	}
	var SomaNotas float64
	for _, nota := range NotasAlunos {
		SomaNotas += nota
	}
	Media := SomaNotas / float64(len(NotasAlunos))
	var menorNota float64 = 11
	var maiorNota float64 = -1
	for _, nota := range NotasAlunos {

		if nota > maiorNota {
			maiorNota = nota
		}
		if nota < menorNota {
			menorNota = nota
		}
	}

	fmt.Printf("A maior nota é: %.2f\n", maiorNota)
	fmt.Printf("A menor nota é: %.2f\n", menorNota)
	fmt.Printf("A média das notas é\n: %.2f", Media)
	
	NomeAlunosAprov := make([]string,0)

	for nome, nota := range NotasAlunos {
		if nota > 7.5 {
			fmt.Printf("O aluno %s foi aprovado com a nota %.2f\n", nome, nota)
			NomeAlunosAprov = append(NomeAlunosAprov, nome)
		}
	}

	fmt.Printf("Os alunos aprovados foram: %v \n", NomeAlunosAprov)

}
