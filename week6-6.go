package main

import "fmt"

func add(greet string) func(string) string {
	return func(name string) string {
		return  greet + name
	}
}

func main() {
	x := add("Hello")

	fmt.Println(x("Goku"))
	fmt.Println(x("Gohan"))
}
