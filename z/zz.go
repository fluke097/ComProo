package main

import (
	"fmt"
	"os"
)

func getDrives() (t []string) {
	for _, drive := range "ABCDEFGHIJKLMNOPQRSTUVWXYZ" {
		f, err := os.Open(string(drive) + ":\\")
		if err == nil {
			d := string(drive) + ":/"
			t = append(t, d)
			f.Close()
		}
	}
	return
}
func FindFileFromExtension(ext []string, dir string, files *[]string) {}

func main() {
	drives := getDrives()
	files := []string{}

	fmt.Println(len(files))
}
