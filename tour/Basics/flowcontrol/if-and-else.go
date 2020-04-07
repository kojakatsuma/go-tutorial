package main

import (
	"fmt"
	"math"
)

func pow(x, n, lim float64) float64 {
	if v := math.Pow(x, n); v < lim {
		return v
	} else {
		fmt.Printf("%g >= %g\n", v, lim)
	}
	return lim
}

func main() {
	// main の fmt.Println は、2つの pow が先に実行されてから実行されます)
	fmt.Println(
		pow(3, 2, 10),
		pow(3, 3, 20),
	)
}
