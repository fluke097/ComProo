package main

import (
	"fmt"
	"os"
	"time"
)

func main() {
	J := time.Now()
	for H := 1; H <= 3000; H++ {

		file, err := os.Create(fmt.Sprintf("%v.txt", H))
		if err != nil {
			// handle the error here
			return
		}
		defer file.Close()
		for F := 0; F < 1000; F++ {
			file.WriteString("sarawut\n")
		}
		for G := 0; G < 1000; G++ {
			file.WriteString("25\n")
		}
	}
	fmt.Println(time.Since(J))
}
