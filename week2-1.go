package main

import "fmt"

func main() {
	fmt.Print(`\n \t t Backticks`)
	fmt.Print("\n \tDouble quote")
	fmt.Println()
	fmt.Println(`\n \t Backticks`[0])
	fmt.Println(len(`\n \tBackticks`))
	fmt.Println("Hello" + "World")
}
