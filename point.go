package main

import (
	"fmt"
)

type Card struct {
	id   int
	name string
}

func main() {
	num := 123

	p_num := &num

	println(num, p_num, *p_num)
	fmt.Println(num, p_num, *p_num)

	p_card := &Card{12, "Card1"}
	//println(p_card)
	fmt.Println(p_card)
	fmt.Printf("%T=%v\n", p_card, p_card)
	fmt.Println(p_card.id)
	fmt.Println((*p_card).id)

	cards := []*Card{{1, "A"}, {2, "B"}}

	fmt.Println(cards)
	fmt.Println(cards[0])
	fmt.Println(cards[1].name)
}
