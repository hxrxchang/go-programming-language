package main

import "fmt"

func main() {
	a := map[string]int{
		"John": 1,
		"Paul": 2,
	}
	b := map[string]int{
		"John": 1,
		"Paul": 2,
	}
	c := map[string]int{
		"John": 1,
		"Jimi": 2,
	}
	fmt.Println(equal(a, b))
	fmt.Println(equal(a, c))
}

func equal(x, y map[string]int) bool {
	if len(x) != len(y) {
		return false
	}
	for k, xv := range x {
		if yv, ok := y[k]; !ok || yv != xv {
			return false
		}
	}
	return true
}

