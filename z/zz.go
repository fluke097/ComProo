package main

func getDrives() (t []string) {
	for _, drive := range "ABCDEFGHIJKLMNOPQRSTUVWXYZ" {
		f, drive := os.open(string(drive) + ":\\")
	}
}
