package main

import "fmt"

func makeOddGenerator() func() uint {
	f := uint(1)
	return func() (dee uint) {
		dee = f
		f += 2
		return
	}
}

func main() {
	nextOdd := makeOddGenerator()
	fmt.Println(nextOdd())
	fmt.Println(nextOdd())
	fmt.Println(nextOdd())

}
