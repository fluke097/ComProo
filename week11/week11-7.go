package main

import "fmt"

func fibonacci() func() int {
	temp := [2]int{0, 0}
	return func() int {
		if temp[1] == 0 {
			temp[1] = 1
			return 0
		}
		temp[0], temp[1] = temp[1], temp[0]+temp[1]
		return temp[0]
	}
}

func main() {
	f := fibonacci()
	for i := 0; i < 10; i++ {
		fmt.Println(f())
	}
}
