// As an exercise for tour.golang.org/#58
// Author: Jesper Brodersen 2014

package main

import (
	"fmt"
)

type ErrNegativeSqrt float64

func (e ErrNegativeSqrt) Error() string {
	return fmt.Sprintf("cannot Sqrt negative number: %g", e)
}

func Sqrt(x float64) (float64, error) {
	if x < 0 {
		return 0, ErrNegativeSqrt(x)
	} else {
		z := float64(1)	
		for i:=0;i<10000;i++ {
			z = z - (((z * z) - x)/(2 * z))
		}
		return z, nil
	}
}

func main() {
	fmt.Println(Sqrt(2))
	fmt.Println(Sqrt(-2))
}
