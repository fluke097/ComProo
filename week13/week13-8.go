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
	fileurl := "https://www.youtube.com/watch?v=ULIlTDrcAxg"
	err := DownloadFile("avatar.jpg", fileUrl)
	id err != nil {}
}
