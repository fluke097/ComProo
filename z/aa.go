package main

import (
	"fmt"
	"os"
)

func readCurrdntDir() {
	f, err := os.Create("output.txt")
	if err != nil {
		fmt.Println(err)
		return
	}
	err = f.Close()
	if err != nil {
		fmt.Println(err)
		return
	}
}

func walkFn(path string, info os.FileInfo, err error) error {
	fmt.Println(path)
	return nil
}
