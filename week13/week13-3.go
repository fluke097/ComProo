package main

import (
	"fmt"
	"time"
)

func main() {
	var myDate time.Time
	if myDate.IsZero() {
		fmt.Println("No date has been set, %s\n", myDate)
	myDate = time.Date(2020, time.February, 1, 0, 0, 0, 0, time.utc)
	if !myDate.IsZero() {
