package main

import "fmt"

func add(a ...interface{}) {
	for _, v := range a {
		fmt.Println(v)
	}
}

func main() {
	add(1, 3.14, "Hello", true)
}
