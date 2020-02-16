package main

	"log"

func main() {
	content, err := ioutil.ReadFile("golangcode.txt")
	if err != nil {
		log.Fatal(err)
}
