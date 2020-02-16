package main

import (
	"fmt"
	"sync"
	"time"
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

func longConcurrentProcess(sleep int, wg *sync.WaitGroup) {
	defer wg.Done()
	time.Sleep(time.Duration(sleep) * time.Second)
	fmt.Println("Sleeping for", sleep, "secounds")
}
