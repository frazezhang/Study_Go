package main

import (
	"fmt"
)

type Card struct {
	id   int
	name string
}

func main() {
	array := [...]Card{{1, "A"}, {2, "B"}}
	slice := []Card{{1, "A"}, {2, "B"}}

	slice_copy := slice // slice_copy point to the slice
	array_copy := array // array_copy is another array

	fmt.Println("array:", array, array_copy)
	fmt.Println("slice:", slice, slice_copy)

	slice_copy[0].id = 3
	array_copy[0].id = 3
	fmt.Println("array:", array, array_copy)
	fmt.Println("slice:", slice, slice_copy)
}
