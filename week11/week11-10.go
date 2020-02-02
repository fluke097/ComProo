package main

import (
	"fmt"
	"math"
)

func main() {
	e := calculateEccentricity(9, 2)
	fmt.Println(e)

}
func calculateEccentricity(a, b float64) float64 {
	return (math.Sqrt(math.Pow(a, 2) - math.Pow(b, 2))) / a
}
