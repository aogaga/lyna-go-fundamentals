package main

import (
	"fmt"
)

func main()  {
	var colors [3]string
	colors[0] = "red"
	colors[1] = "green"
	colors[2] = "blue"
	fmt.Println(colors)
	fmt.Println(colors[1])


	var numbers = [5]int{5,3,1,2,4}
	fmt.Println(numbers)


	fmt.Println("number of colors: ", len(colors))
	fmt.Println("Number of numers", len(numbers))
}