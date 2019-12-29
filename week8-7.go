package main

import "fmt"

type I interface{}

func desc(i I) {
	fmt.Printf("%v , %T \n", i, i)
}

func main() {
	var i I
	i = 2.880
	desc(i)
	i = "Hello saraut"
	desc(i)
}
