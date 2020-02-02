package main

import (
	"fmt"

	"github.com/goyod/labs/fizzbuzz"
)

func main() {
	for i := 1; i <= 100; i++ {
		fmt.Print(fizzbuzz.FizzBuzz(i) + ",")
	}
}
