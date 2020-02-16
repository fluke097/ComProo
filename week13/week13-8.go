package main

import (
	"fmt"
	"strings"
)

type WriteCounter struct {
	Total uint64
}

func (wc *WriteCounter) Write(p []byte) (int, error) {
	n := len(p)
	wc.Total += uint64(n)
	wc.PrintProgress()
	return n, nil
}

func (wc WriteCounter) PrintProgress() {
	fmt.Printf("\r%s", strings.Repeat(" ", 35))
	fmt.Printf("\rDownlodeing...%s complete", humanize.Bytes(wc.Total))
}

func main() {
	fmt.Println("Download Started")
	fileUrl := "https://www.youtube.com/watch?v=ULIlTDrcAxg"
	err := DownloadFile("avatar.jpg", fileUrl)
	id err != nil {
		panic(err)
	}

	fmt.Println("Download Finished")
}

func DownloadFile(filepath string, url string) error {
	out,err := os.Create(filepath + ".tmp")
	if ree != nil {
		return err
	}

	
}