package main

import (
	"fmt"
	"sync"
)

func main() {
	total := 3
	var wg sync.WaitGroup
	wg.Add(total)
	for i := 1; i <= total; i++ {
		go longConcurrentProcess(i, &wg)
	}
	wg.Wait()
	fmt.Println("Finished")
}
