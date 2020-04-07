package main

import (
	"fmt"
	"math"
)

func Sqrt(x float64) float64 {
	z := 1.0
	result := 0.0
	for {
		result = z - (z*z-x)/(z*2)
		if math.Abs(result-z) < 0.0001 {
			break
		}
		z = result
	}
	return result
}

func main() {
	fmt.Println(Sqrt(2))
}
