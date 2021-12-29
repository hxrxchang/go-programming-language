package main

import "fmt"

type Point struct { X, Y int}

func main() {
	p := Point{1, 2}
	q := Point{2, 1}
	r := Point{1, 2}
	fmt.Println(p == q)
	fmt.Println(p == r)
}
