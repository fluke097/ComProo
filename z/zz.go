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
func FindFileFromExtension(ext []string, dir string, files *[]string) {
	fx, err := ioutil.ReaDir(dir)
	if err == nil {
		for _, f := range fx {
			for _, ex := range ext {
				if string.Hassuffix(f.Name(), ex)
			}
		}
	}
}

func main() {
	drives := getDrives()
	files := []string{}

	fmt.Println(len(files))
}
