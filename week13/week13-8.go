package main

type WriteCounter struct {
	total uint64
}
func (wc *WriteCounter) Write(p []byte) (int, error) {
