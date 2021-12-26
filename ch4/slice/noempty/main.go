package main

import "fmt"

func main() {
	data := []string{"one", "", "three"}
	data = noempty(data)
	fmt.Printf("%q\n", data)
}

func noempty(strings []string) []string {
	out := strings[:0]
	for _, s := range strings {
		if s != "" {
			out = append(out, s)
		}
	}
	return out
}
