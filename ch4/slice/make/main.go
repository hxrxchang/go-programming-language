package main

import "fmt"

func main() {
	s := make([]string, 5)
	for i, v := range(s) {
		fmt.Println(i, v, v == "")
	}
}
