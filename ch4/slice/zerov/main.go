package main

import "fmt"

func main() {
	var s []int
	fmt.Println(s, len(s), s == nil)
	s = nil
	fmt.Println(s, len(s), s == nil)
	s = []int(nil)
	fmt.Println(s, len(s), s == nil)
	s = []int{}
	fmt.Println(s, len(s), s == nil)
}
