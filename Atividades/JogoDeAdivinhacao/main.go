package main

import (
	"bufio"
	"fmt"
	"math/rand/v2"
	"os"
	"strconv"
	"strings"
)

func main() {
	fmt.Println("Jogo da Adivinhaçao")
	fmt.Println("Tente adivinhar o número entre 1 e 100")

	scanner := bufio.NewScanner(os.Stdin)
	chutes := [10]int64{}
	NumeroAlea := rand.Int64N(101)
	for i := range chutes {
		fmt.Println("Digite seu chute")
		scanner.Scan()
		chute := strings.TrimSpace(scanner.Text())
		chuteInt, err := strconv.ParseInt(chute, 10, 64)
		if err != nil {
			fmt.Println("Erro ao converter o número")
			return
		}

		switch {
		case chuteInt < NumeroAlea:
			fmt.Println("O nome sorteado é maior que", chuteInt)
		case chuteInt > NumeroAlea:
			fmt.Println("O numero sorteado é menor que", chuteInt)
		default:
			fmt.Printf(
				"Parabens, voce acertou o numero, que era: %d\n"+
					"Voce acertou em %d tentativas. \n"+
					"Essas foram as tentativas: %v\n", NumeroAlea,i+1 ,chutes [:i],
			)
			return
		}
		chutes[i] = chuteInt
	}
	fmt.Printf(
		"Infelizmente, voce nao acertou o numero, que era: %d\n"+
			"Voce teve 10 tentativas. \n"+
			"Essas foram as tentativas: %v\n", NumeroAlea, chutes,
	)
}
