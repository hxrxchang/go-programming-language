package main

import "fmt"

var a [3]int
var b [2]int = [2]int{1, 2}
var c [2]int = [...]int{1, 2}
var d [2]int = [2]int{1, 3}

func main() {
	fmt.Println(a[0])
	fmt.Println(a[len(a) - 1])

	for i, v := range a {
		fmt.Printf("%d %d\n", i, v)
	}

	fmt.Println(b == c, b == d, c == d)
}
