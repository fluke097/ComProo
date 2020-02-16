package mainpackage main
func main() {
	myList := map[String]string{
		"dog":       "woof",
		"cat":       "meow",
		"hedgenhog": "sniff",

	for animal, noise := range myList {
		fmt.Println("The", animal, "went", noise)
