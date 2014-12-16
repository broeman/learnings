// Newton's approximation of square root
// Author: Jesper Brodersen 2014


package main

import (
	"fmt"
	"math"
)

// As an exercise for tour.golang.org/#25
func SimpleSqrt(x float64) float64 {
	z := float64(1)
	for i:=0;i<10000;i++ {
		z = z - (((z * z) - x)/(2 * z))
	}
	return z
}


// Using http://en.wikipedia.org/wiki/Newton's_method as inspiration
func f(x, y float64) float64 {
	return (x * x) - y
}

func fprime(x float64) float64 {
	return 2 * x
}

func WikiSqrt(x float64) float64 {
	x0 := float64(1)
	const tolerance = 1e-7
	const epsilon = 1e-15
	x1 := x0

	for i:=0; i < 20; i++ {
		y := f(x0, x)
		yprime := fprime(x0)

		if math.Abs(yprime) < epsilon {
			fmt.Println("Denominator is too small")
			break
		}

		x1 := x0 - y / yprime

		if math.Abs(x1 - x0) / math.Abs(x1) < tolerance {
			return x1
		}

		x0 = x1		
	}
	return x1
}

func main() {
	fmt.Println(SimpleSqrt(2))
	fmt.Println(WikiSqrt(2))
	fmt.Println(math.Sqrt(2))
	fmt.Println("")
	fmt.Println(SimpleSqrt(4))
	fmt.Println(WikiSqrt(4))
	fmt.Println(math.Sqrt(4))
	fmt.Println("")
	fmt.Println(SimpleSqrt(8))
	fmt.Println(WikiSqrt(8))
	fmt.Println(math.Sqrt(8))
}
