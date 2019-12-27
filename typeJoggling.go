package main

import (
	"fmt"
)

func main()  {

	var int1 int = 5;
	var float1 float64 = 42;
	sum := float64(int1) * float1
	fmt.Printf("Sum: %v type: %T", sum , sum)
}