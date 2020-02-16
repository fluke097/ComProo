package mainpackage main

import "fmt"

func main() {
	myList := map[string]string{
		"dog":       "woof",
		"cat":       "meow",
		"hedgenhog": "sniff",
	}

	for animal, noise := range myList {
		fmt.Println("The", animal, "went", noise)
	}
}
