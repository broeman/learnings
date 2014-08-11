package main

import (
	"fmt"
	"github.com/broeman/testlib"
)
	
func main() {
	t := Testme{"the test worked"}
	fmt.Println(t.output)
}