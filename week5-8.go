package main

import "fmt"

func main() {
	humans := []string{"Bulma", "Chi-Chi", "Videl"}
	names := []string{"Goku", "Gohan"}
	names = append(names, humans...)
	fmt.Println(names)
}
