package main

import "fmt"

func main()  {
	// var card string ="ace of assaf"
	card := newCard()
	// := is only used when variable is declared for first time
	// when variable is updated only = is used
	fmt.Println(card)
	var in int = 43
	// in := 45
	fmt.Println(in)
	
}
func newCard() string{
	return "five of diamonds"
}

