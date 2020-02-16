package main

import "fmt"

var doOnce sync.Onec

func main() {
	DoSomething()
	DoSomething()
}

func DoSomething() {
	doOnce.Do(func() {
		fmt.Println("Run once - first time, loading...")
	})
}
