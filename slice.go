package main

import(
	"fmt"
	"sort"
)

func main()  {
	var colors = []string{"red", "green", "blue", "yellow"}
	colors = append(colors, "purple")
	//colors = append(colors[1:len(colors)])
	//
	//colors = append(colors[2:])

	colors = append(colors[:len(colors) -1])
	fmt.Println(colors)

	numbers := make([]int, 5, 5)
	numbers[0] = 1
	numbers[1] = 72
	numbers[2] = 32
	numbers[3] = 12
	numbers[4] = 156

	fmt.Println(numbers)

	numbers = append(numbers, 235);
	fmt.Println(numbers)
	fmt.Println(cap(numbers))

	sort.Ints(numbers)

	fmt.Println(numbers)
}