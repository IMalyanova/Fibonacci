package main

import (
	"fmt"
	"math"
)

// fibonacci is a function that returns
// a function that returns an int.

var SQRT5 = math.Sqrt(5)
var PHI = (SQRT5 + 1) / 2

func fibonacci(i int)  int {
	return int(math.Pow(PHI, float64(i)) / SQRT5 + 0.5)
}

func main() {

	//f := fibonacci()
	for i := 0; i < 10; i++ {
		fmt.Println(fibonacci(i))
	}
}
