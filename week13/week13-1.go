package main
func readword(ch chan string) {}
func readword(ch chan string) {
	fmt.Println("Type a word, then hit Enter.")
	var word string
	fmt.Scanf("%s", &word)
	ch <- word
}
func timeout(t chan bool) {
	time.Sleep(5 * time.Second)
	t <- false
