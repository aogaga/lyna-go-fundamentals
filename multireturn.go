package main

import(
	"fmt"
)

func main()  {

	n1, l1 := FullName("Zaphood", "beeblebrox")
	fmt.Println(n1)
	fmt.Println("\n")
	fmt.Println(l1);


	n2, l2 := FullNameNakedReturn("Arthur", "Dent")
	fmt.Printf("Fullname: %v, number of chars: %v", n2, l2)
}

func FullName(f, l string)(string, int)  {
	full := f + " "+ l
	length := len(full)
	return full, length
}

func FullNameNakedReturn(f, l string)(full string, length int)  {
	full = f + " "+ l
	length = len(full)
	return
}