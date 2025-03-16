package main

import (
	"fmt"
)

func main() {
	najmniejsza := 0
	i := 1
	for (i != 0){
		fmt.Print("Wpisz liczbe calkowita: ")
		fmt.Scan(&i)

		if i < najmniejsza {
			najmniejsza = i
		}
	}
	fmt.Println("Najmniejsza wprowadzona liczba to:", najmniejsza)
}