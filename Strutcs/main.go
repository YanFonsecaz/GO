package main

import "fmt"

func main () {
	p := Pessoas{
		Nome: "Yan",
		Idade: 25,
		Email: "yan@example.com",
	}
	p.ExibirDados()
	fmt.Println("É maior de idade?", p.EhMaiorDeIdade())

	r := Retangulo{
		Altura: 10,
		Largura: 5,
	}
	fmt.Println("A área do retângulo é:", r.Area())
	fmt.Println("O perímetro do retângulo é:", r.Perimetro())
	fmt.Println("O retângulo é um quadrado?", r.EhQuadrado())
	r.ExibirValoresRetangulo()
}

type Pessoas struct {
	Nome string
	Idade uint64
	Email string
}

func (p Pessoas) EhMaiorDeIdade() bool {
	return p.Idade >= 18
}

func (p Pessoas) ExibirDados() {
	fmt.Println("Nome:", p.Nome)
	fmt.Println("Idade:", p.Idade)
	fmt.Println("Email:", p.Email)
}

type Retangulo struct{
	Altura float64
	Largura float64
}

func (r Retangulo) Area () float64{
	return r.Altura * r.Largura
}

func (r Retangulo) Perimetro () float64{
	return 2 * (r.Altura + r.Largura)
}

func (r Retangulo) EhQuadrado() bool{
	return r.Altura == r.Largura
}

func (r Retangulo) ExibirValoresRetangulo() {
	fmt.Println("Altura:", r.Altura)
	fmt.Println("Largura:", r.Largura)
}
