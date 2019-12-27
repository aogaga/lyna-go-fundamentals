package main

import(
	"fmt"
)

func main()  {
	var p *int

	if p != nil{
		fmt.Println("value of : ", *p)
	}else{
		fmt.Println("p is nill")
	}

	var v int = 42
	p = &v


	if p != nil{
		fmt.Println("value of : ", *p)
	}else{
		fmt.Println("p is nill")
	}

	var value1 float64 = 42.13
	pointer1 := &value1
	fmt.Println("value 1 is : ", *pointer1)

	*pointer1 = *pointer1 /31
	fmt.Println("new value of Pointer is ", *pointer1)
}
