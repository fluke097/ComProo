package main

import (
	"fmt"
	"sort"
	"strconv"
)

func main() {
	myList := []string{"1", "3", "2", "9", "6", "5", "4", "7", "10", "8"}

	fmt.Printf("Before: %v\n", myList)
	sort.Slice(myList, func(i, j int) bool {
		numA, _ := strconv.Atoi(myList[i])
		numB, _ := strconv.Atoi(myList[j])
		return numA < numB
	})

	fmt.Printf("After: %v\n", myList)
}
