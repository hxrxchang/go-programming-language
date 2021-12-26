package main

import "fmt"

func main() {
	a := []string {"a", "b", "c"}
	b := []string {"a", "b", "c"}
	c := []string {"a", "b", "d"}
	fmt.Println(equal(a, b))
	fmt.Println(equal(a, c))
}

func equal(x, y []string) bool {
	if len(x) != len(y) {
		return false
	}
	for i := range x {
		if x[i] != y[i] {
			return false
		}
	}
	return true
}
