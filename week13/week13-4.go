package main

import "fmt"

func main() {
	result := ReturnData()
	myInt := result.(int)
	myInt += 5
	fmt.Println(myInt)
}

func ReturnData() interface{} {
	return 5
}
