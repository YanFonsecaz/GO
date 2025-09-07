package main

import "fmt"

func main() {
	Pessoas := map[string]int{
		"Yan": 25,
		"João": 30,
		"Maria": 20,
		"Pedro": 35,
		"Paulo": 40,
		"Paula": 45,
		"Jose": 60,
		"Joaquina": 65,
		"Joaquim": 70,
		"Joana": 75,
	}
	Crianca := make(map[string]int)
	Adolescente := make(map[string]int)
	Adulto := make(map[string]int)
	Idoso := make(map[string]int)
	for nome, idade := range Pessoas {
		if idade >= 0 && idade <= 12 {
			Crianca[nome] = idade
		} else if idade >= 13 && idade <= 18 {
			Adolescente[nome] = idade
		} else if idade >= 19 && idade <= 65 {
			Adulto[nome] = idade
		} else {
			Idoso[nome] = idade
		}
	}

	fmt.Println("FAIXA ETARIA")

	// Verificação e exibição de Crianças
	if len(Crianca) == 0 {
		fmt.Println("Nenhuma Criança")
	} else {
		for nome, idade := range Crianca {
			fmt.Println("Criança:", nome, idade)
		}
	}

	// Verificação e exibição de Adolescentes
	if len(Adolescente) == 0 {
		fmt.Println("Nenhum Adolescente")
	} else {
		for nome, idade := range Adolescente {
			fmt.Println("Adolescente:", nome, idade)
		}
	}

	// Verificação e exibição de Adultos
	if len(Adulto) == 0 {
		fmt.Println("Nenhum Adulto")
	} else {
		for nome, idade := range Adulto {
			fmt.Println("Adulto:", nome, idade)
		}
	}

	// Verificação e exibição de Idosos
	if len(Idoso) == 0 {
		fmt.Println("Nenhum Idoso")
	} else {
		for nome, idade := range Idoso {
			fmt.Println("Idoso:", nome, idade)
		}
	}
}