package main

import (
	"awesomeProject/pokemon"
	"fmt"
)

func main() {
	pickachu := pokemon.Pokemon{
		ID:   1,
		NOM:  "PICKACHU",
		TYPE: "ELE",
	}
	fmt.Println("nom :\n", pickachu.NOM, "type :", pickachu.TYPE)
}
