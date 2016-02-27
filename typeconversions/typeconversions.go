package main

import (
	"fmt"
	"math"
)

func main() {
	var x, y int = 3, 4
	// Assignment between items of different types requires an explicit
	// conversion. 
	var f float64 = math.Sqrt(float64(x*x + y*y))
	var z uint = uint(f)
	fmt.Println(x, y, z)
}
