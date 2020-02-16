package main

import "sync"

func main() {
	total := 3
	var wg sync.WaitGroup
	wg.Add(total)
}
