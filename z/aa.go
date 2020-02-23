package main

import (
	"fmt"
	"os"
)

func readCurrdntDir() {
	f, err := os.Create("output.txt")
	if err != nil {
		fmt.Println(err)
		return
	}
}
