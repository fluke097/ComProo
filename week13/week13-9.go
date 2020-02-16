package main
import (
	"fmt"
	"os"
)
func main() {
	fmt.Println(os.Getenv("VISUAL"))
	os.Setenv("Site", "GoLangCode")
	fmt.Println(os.Getenv("site"))
	fmt.Println(os.Getenv("missing"))
}
