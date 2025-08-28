package main

import "fmt"
func main() {
	for i := 1 ; i <= 10 ; i++{
		fmt.Println(i)
	}
	for i := 1; i<= 10 ; i++ {
		Tabuada := 7
		Resultado := Tabuada * i
		fmt.Printf("%d x %d = %d\n",Tabuada,i,Resultado)
	}
	for i := 10 ; i >= 1 ; i-- {
		fmt.Printf("%d\n",i)
	}
	var Total int
	for i := 1 ; i<= 100 ; i++ {
		Total += i 
	}
	fmt.Printf("%d\n",Total)
	for i := 1 ; i<= 20; i++ {
		if i % 2 == 0 {
			fmt.Printf("%d é par\n",i)
		}
	}
	for i := 1; i <= 20; i++ {
		if i == 15 {
			fmt.Printf("Encontrei o %d \n",i)
			break
		}
	}

}