package main

import (
	"fmt"
	"os"
	"strings"
	"bufio"
)

func main() {
	Contatos := [][2]string{{"Yan", "449399434"}, {"Teste", "449399434"}, {"Joao", "443235532"}}
	scanner := bufio.NewScanner(os.Stdin)
	for i, contato := range Contatos {
		fmt.Printf("%d : %s - %s \n",i,contato[0],contato[1])
	}
	fmt.Println("Digite o nome do contato que deseja buscar")
	scanner.Scan()
	nome := strings.ToLower(strings.TrimSpace(scanner.Text()))
	
	encontrado := false
	for _, contato := range Contatos {
		if strings.ToLower(contato[0]) == nome {
			fmt.Printf("Contato encontrado: %s - %s \n", contato[0],contato[1])
			encontrado = true
			break
		}
	}
	if !encontrado {
		fmt.Println("Contato não encontrado")
	}
}
