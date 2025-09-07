package internal


type SistemaProdutos struct {
	Produtos []Produto
}

type Produto struct {
	ID        uint64
	Nome      string
	Preco     float64
	Estoque   uint64
	Categoria string
}

// Adiciona Novos Produtos
func (s *SistemaProdutos) AdicionarProduto(produto Produto) {
	s.Produtos = append(s.Produtos, produto)
}

// Aplica Desconto no Produto
func (p *Produto) AplicarDesconto(desconto float64) {
	p.Preco = p.Preco - (p.Preco * (desconto / 100))
}

// Informa o valor total do estoque
func (p Produto) ValorTotalEstoque() float64 {
	return p.Preco * float64(p.Estoque)
}

// Verifica se o produto está em falta
func (p Produto) EmFalta() string {
	if p.Estoque < 5 {
		return "Fora do Estoque"
	}
	return "Em Estoque"
}