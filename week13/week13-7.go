package main

import (
	"fmt"
	"sort"
)

func main() {
	myList := []string{"1", "2", "3", "4", "5", "6", "7", "8", "9", "10"}

	fmt.Printf("Before: %v\n", myList)
	sort.Slice(myList, func(i, j int) bool {

	})
}
