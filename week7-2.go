package main

import "fmt"

func handlePanic() {
	r := recover()
	if r == "to much" {
		fmt.Println("your number out of range")
		main()
	}
}

func main() {
	defer handlePanic()
	var i int
	fmt.Print("type number :")
	_, e := fmt.Scan(&i)
	if e != nil {
		panic("to much")
	}
	fmt.Println("your number :", i)
}
