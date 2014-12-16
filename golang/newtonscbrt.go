package main

import (
    "fmt"
    "math/cmplx"
)

func Cbrt(x complex128) complex128 {
    x0 := x
    const epsilon = 1e-15
    for {
        x1 := x0 - ((x0 * x0 * x0) - x) / (3 * (x0 * x0))
        if cmplx.Abs(x1 - x0) < epsilon {
        	return x0
        }
        x0 = x1
    }   
    return x0
}

func main() {
    fmt.Println(Cbrt(2))
    fmt.Println(cmplx.Pow(2,1.00/3.00))
}