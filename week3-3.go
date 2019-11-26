package main

import "fmt"

func main() {
	n, e := fmt.Println("Hello", "World", 123, 456)
	fmt.Println("number of bytes writen :", n)
	fmt.Println("write error encountered :", e)
}
