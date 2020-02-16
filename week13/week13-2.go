package main

import (
	"fmt"
	"io/ioutil"
	"log"
)

func main() {
	content, err := ioutil.ReadFile("golangcode.txt")
	if err != nil {
		log.Fatal(err)
	}

	text := string(content)
	fmt.Println(text)
}
