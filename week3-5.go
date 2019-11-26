package main

import "fmt"

func main() {
	fmt.Print("input : ")
	var name string
	var age int
	var height float32
	var weigth float32
	n, err := fmt.Scan(&name, &age, &weigth, &height)
	fmt.Println(name, age, weigth, height)
	fmt.Println(`number of argument `, n)
	fmt.Println(`error  `, err)
}
