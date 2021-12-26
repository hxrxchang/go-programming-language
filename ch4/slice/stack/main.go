package main

import "fmt"

func main() {
	stack := []int{}
	stack = append(stack, 1, 2, 3)
	fmt.Println(stack)
	stack = stack[:len(stack)-1]
	fmt.Println(stack)
}
