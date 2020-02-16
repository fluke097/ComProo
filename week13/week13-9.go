package main
func main() {
	fmt.Println(os.Getenv("VISUAL"))
	os.Setenv("Site", "GoLangCode")
	fmt.Println(os.Getenv("site"))
}
