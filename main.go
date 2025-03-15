package main

import (
	"fmt"
)

func main() {
	najmniejsza := 0
	for {
		var i int
		fmt.Print("Wpisz liczbe calkowita: ")
		fmt.Scan(&i)

		if i == 0 {
			break
		}
		if najmniejsza == 0 || i < najmniejsza {
			najmniejsza = i
		}
	}
	fmt.Println("Najmniejsza wprowadzona liczba to:", najmniejsza)
}
