package main

import "fmt"

func main() {
	i := 0
	toTen(&i)
	fmt.Println(i)
}

func toTen(i *int) {
	if *i >= 10 {
		return
	}
	*i = *i + 1
	toTen(i)
}
