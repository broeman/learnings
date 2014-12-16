package main

import (
	"fmt"
	"github.com/broeman/testlib"
)

func main() {
	t := testlib.Testme{}
	t.SetInput("The setter worked")
	output := t.Input()
	fmt.Println(output)
}
