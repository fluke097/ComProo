package main

var doOnce sync.Onec

func main() {
	DoSomething()
	DoSomething()
}
