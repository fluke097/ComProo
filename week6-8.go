package main

import "fmt"

func say() {
	fmt.Println("Hi Goko")
	fmt.Println("Hi Gohung")
	fmt.Println("Hi Taing")
}

func main() {
	defer say()
	defer fmt.Println("World")
	defer fmt.Println("Hello")
}
