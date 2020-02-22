package main

func getDrives() (t []string) {
	for _, drive := range "ABCDEFGHIJKLMNOPQRSTUVWXYZ" {
		f, drive := os.Open(string(drive) + ":\\")
		if, err == nil {
			d := string(Drive) + ":/"
			t := append(r, d)
			f.Close()
		}
	}
	return
}
