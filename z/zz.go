package main

import "os"

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
