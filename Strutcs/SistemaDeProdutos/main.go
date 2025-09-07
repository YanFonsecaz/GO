package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
	"v11/internal"
)

func main() {

	scanner := bufio.NewScanner(os.Stdin)
	sistema := internal.SistemaProdutos{}

	fmt.Printf("Sistema de Produtos\n")
	fmt.Printf("------------------\n")
	fmt.Printf(
		"Escolha uma das opções: \n" +
			"1 - Adicionar Produto\n" +
			"2 - Aplicar Desconto\n" +
			"3 - Verificar Estoque\n" +
			"4 - Sair\n")
	scanner.Scan()
	opcao, _ := strconv.ParseInt(scanner.Text(), 10, 64)
	switch opcao {
	case 1:

		fmt.Printf("Adicionar Produto\n")
		fmt.Printf("------------------\n")

		// Adicionar ID do Produto
		fmt.Printf("ID: \n")
		scanner.Scan()
		id, err := strconv.ParseInt(scanner.Text(), 10, 64)
		if err != nil {
			fmt.Printf("Error Valor do ID invalido: %v\n", err)
			return
		}

		// Adicionar Nome do Produto
		fmt.Printf("Nome do Produto: \n")
		scanner.Scan()
		nome := strings.ToLower(strings.TrimSpace(scanner.Text()))
		if nome == "" {
			fmt.Printf("Error Nome do Produto invalido\n")
			return
		}

		// Adicionar Preço do Produto
		fmt.Printf("Preço: \n")
		scanner.Scan()
		preco, err := strconv.ParseFloat(scanner.Text(), 64)
		if err != nil {
			fmt.Printf("Erro Valor do Preço invalido: %v\n", err)
			return
		}

		// Adicionar Estoque do Produto
		fmt.Printf("Estoque: \n")
		scanner.Scan()
		estoque, err := strconv.ParseInt(scanner.Text(), 10, 64)
		if err != nil {
			fmt.Printf("Error Valor do Estoque invalido: %v\n", err)
			return
		}
		if estoque <= 0 {
			fmt.Printf("Error Valor do Estoque invalido\n")
			return
		}

		// Adicionar Categoria do Produto
		fmt.Printf("Categoria: \n")
		scanner.Scan()
		categoria := scanner.Text()
		if categoria == "" {
			fmt.Printf("Error Categoria do Produto invalido\n")
			return
		}

		// Adicionar Produto
		produto := internal.Produto{
			ID:        uint64(id),
			Nome:      nome,
			Preco:     preco,
			Estoque:   uint64(estoque),
			Categoria: categoria,
		}
		sistema.AdicionarProduto(produto)
		fmt.Printf("Produto Adicionado com Sucesso\n")
	case 2:
		fmt.Printf("Aplicar Desconto\n")
		fmt.Printf("------------------\n")
		fmt.Printf("Nome do Produto: \n")
		scanner.Scan()
		nome := strings.ToLower(strings.TrimSpace(scanner.Text()))

		var produtoEncontrado *internal.Produto
		for i := range sistema.Produtos {
			if nome == sistema.Produtos[i].Nome {
				produtoEncontrado = &sistema.Produtos[i]
			}
		}

		if produtoEncontrado != nil {
			fmt.Printf("Produto Encontrado\n")
			fmt.Printf("Qual o valor do desconto %%? \n")
			scanner.Scan()
			desconto, err := strconv.ParseFloat(scanner.Text(), 64)
			if err != nil {
				fmt.Printf("Error Valor do Desconto invalido: %v\n", err)
				return
			}
			produtoEncontrado.AplicarDesconto(desconto)
			fmt.Printf("Preço com Desconto: %v\n", produtoEncontrado.Preco)
			return
		} else {
			fmt.Printf("Produto nao encontrado")
		}
	case 3:
		fmt.Printf("Verificar Estoque\n")
		fmt.Printf("------------------\n")
		
	}
}
