package main

import "fmt"

type Point struct { X, Y float64 }

func main() {
	r := Point{1, 2}
	r.ScaleBy(2)
	fmt.Println(r)
}

func (p *Point) ScaleBy(factor float64) {
	p.X *= factor
	p.Y *= factor
}
