package main

import "fmt"

func main() {
	alPhabets := [4]string{"A", "B", "C", "D"}
	x := alPhabets[:]
	y := alPhabets[2:]
	z := alPhabets[1:]
	fmt.Println(x)
	fmt.Println(y)
	fmt.Println(z)
}
