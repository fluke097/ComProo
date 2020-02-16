package main

import (
	"fmt"
	"sort"
	"strconv"
)

func main() {
	myList := []string{"1", "2", "3", "4", "5", "6", "7", "8", "9", "10"}

	fmt.Printf("Before: %v\n", myList)
	sort.Slice(myList, func(i, j int) bool {

		numA, _ := strconv.Atoi(myList[i])
		numB, _ := strconv.Atoi(myList[j])
	})
}
