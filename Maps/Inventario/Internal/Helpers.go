package internal

import "fmt"

var Produtos = make(map[string]int)

func AdicionarProdutos(NomeProduto string, Quantidade int) {
	if _, ok := Produtos[NomeProduto]; ok {
		fmt.Printf("O produto %s já existe. Atualizando quantidade para %d unidades\n", NomeProduto, Quantidade)
		
	} else {
		fmt.Printf("Produto %s adicionado com %d unidades\n", NomeProduto, Quantidade)
	}
	Produtos[NomeProduto] = Quantidade
}

func RemoverProduto( NomeProduto string) {
	delete(Produtos, NomeProduto)
	fmt.Printf("O produto %s foi removido\n", NomeProduto)
}

func ConsultarEstoque(NomeProduto string) {
	if quantidade, ok := Produtos[NomeProduto]; ok{
		fmt.Printf("O produto %s possui %d unidades em estoque\n", NomeProduto, quantidade)
	} else {
		fmt.Printf("Produto %s não encontrado \n", NomeProduto)
	}
}

func ListarBaixoEstoque() {
	fmt.Printf("Produtos com estoque abaixo de 5 unidades:\n")
	encontrou := false
	for nome, quantidade := range Produtos {
		if quantidade < 5 {
			fmt.Printf("O produto %s possui %d unidades em estoque\n", nome, quantidade)
			encontrou = true
		}
	}
	if !encontrou {
		fmt.Printf("Nenhum produto com estoque abaixo de 5 unidades\n")
	}
}