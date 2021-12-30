package main

import "fmt"

func main() {
	defer func() {
		fmt.Println("End")
		err := recover()
		if err != nil {
			fmt.Printf("Recover!: %v\n", err)
		}
	}()
	fmt.Println("Start")
	panic("Panic")
}
