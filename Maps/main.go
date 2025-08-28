package main

import (
	"fmt"
	"strings"
)

func main() {
	Contatos := make(map[string]string)
	Contatos["Yan"] = "449399434"
	Contatos["Teste"] = "449399434"
	Contatos["Joao"] = "443235532"
	fmt.Printf("%v\n", Contatos)

	nome := "Yan"
	telefone, ok := Contatos[nome]

	if ok {
		fmt.Printf("Contato encontrado %s - %s \n", nome, telefone)
	} else {
		fmt.Printf("Contato não encontrado %s - %v \n", nome, ok)
	}

	frase := "go é legal go é facil"
	fmt.Printf("Frase %s: \n", frase)

	contador := make(map[string]int)

	palavras := strings.Fields(frase)
	fmt.Printf("Palavras %v \n", palavras)
	
	for _, palavra := range palavras {
		contador[palavra]++
	}

	fmt.Println("\n Resultado")
	for palavra, quantidade := range contador {
		fmt.Printf("%s: %d \n", palavra, quantidade)
	}

}
