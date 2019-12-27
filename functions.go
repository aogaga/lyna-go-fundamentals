package main

import (
	"fmt"
)

func main()  {
	dosomething()

	sum := addValues(23, 54)
	fmt.Println(sum)

	sum = addAllValues(12, 54, 79)
	fmt.Println(sum)
}

func dosomething()  {
	fmt.Println("doing something")
}

func addValues( value1 int, value2 int) int  {
	return value1 + value2;
}

func addAllValues( values ...int)int   {
	sum := 0
	for i := range values{
		sum += values[i]
	}
	fmt.Printf("%T\n", values)
	return sum
}