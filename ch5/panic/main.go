package main

import "fmt"

func helloworld() {
	defer fmt.Println("End")

	fmt.Println("Start")
	panic("Panic!")
}

func main() {
	helloworld()
}
