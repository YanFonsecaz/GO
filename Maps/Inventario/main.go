package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
	"v8/Internal"
)

func main() {

	scanner := bufio.NewScanner(os.Stdin)

	for {

			var opcao int64
			fmt.Printf(
				"Qual a opção você quer fazer?\n"+
				"1 - Adicionar produto\n"+
				"2 - Remover produto\n"+
				"3 - Verificar produto\n"+
				"4 - Listar produtos com estoque baixo\n"+
				"5 - Sair\n",
			)
			scanner.Scan()
			opcao, err := strconv.ParseInt((strings.TrimSpace(scanner.Text())),10,64)
			if err != nil {
				fmt.Printf("Erro ao converter a opção para inteiro: %v\n", err)
				return
			}
		
			switch opcao {
			case 1:
				fmt.Printf("Qual o produto você quer adicionar?\n")
				scanner.Scan()
				produto := strings.ToLower((strings.TrimSpace(scanner.Text())))
				fmt.Printf("Qual a quantidade do produto %s você quer adicionar?\n",produto)
				scanner.Scan()
				quantidade, err := strconv.ParseInt((strings.TrimSpace(scanner.Text())),10,64)
				if err!= nil {
					fmt.Printf("Erro ao converter a quantidade para inteiro: %v\n", err)
					return
				}
				internal.AdicionarProdutos(produto,int(quantidade))
			case 2:
				fmt.Printf("Qual produto você deseja remover?\n")
				scanner.Scan()
				produto := strings.ToLower(strings.TrimSpace(scanner.Text()))
				internal.RemoverProduto(produto)
			case 3:
				fmt.Printf("Qual produto você deseja verificar?\n")
				scanner.Scan()
				produto := strings.ToLower(strings.TrimSpace(scanner.Text()))
				internal.ConsultarEstoque(produto)
			case 4:
				internal.ListarBaixoEstoque()
			case 5:
				fmt.Printf("Saindo...\n")
				return
			default:
				fmt.Printf("Opção inválida\n")
			}
	}
	

}