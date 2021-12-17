package main

import "fmt"

const boiliingF = 212.0

func main() {
	var f = boiliingF
	var c = (f - 32) * 5 / 9
	fmt.Printf("boiling point = %gF or %gC\n", f, c)
}
