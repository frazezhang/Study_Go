package main

import (
	"fmt"
)

func main() {
	var i8 int8
	var u16 uint16
	i8 = 127
	u16 = 65535
	fmt.Printf("i8=%v u16=%v\n", i8, u16)

	i8++
	u16++
	fmt.Printf("i8=%v u16=%v\n", i8, u16)

	u16 = uint16(i8)
	fmt.Printf("i8=%v u16=%v\n", i8, u16)
}
