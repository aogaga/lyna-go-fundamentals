package main

import(
	"fmt"
)

func main()  {
	sum := 1
	fmt.Println("SUM", sum)

	colors := []string{"Red", "Green", "Blue"}
	fmt.Println(colors)

	sum = 0

	for i :=0; i < 10 ; i++{
		sum += i
	}

	fmt.Println("SUM", sum)


	for i := 0; i < len(colors); i++{
		fmt.Println(colors[i])
	}


	for i := range colors{
		fmt.Println(colors[i])
	}

	sum = 1

	for sum < 1000 {
		sum += sum;
		if sum > 200{
			goto endofProgram
		}
		if sum > 500{
			break
		}
		fmt.Println("SUM", sum)

	}


	endofProgram : fmt.Println("end of propgram")
}
