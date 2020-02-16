package main

import (
	"fmt"
	"time"
)

func main() {
	var myDate time.Time
	if myDate.IsZero() {
		fmt.Println("No date has been set, %s\n", myDate)
