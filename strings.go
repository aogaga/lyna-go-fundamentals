package main

import(
	"fmt"
	"strings"
)


func main(){
	str1 := "An implicityly typed string"
	fmt.Println(str1)
	fmt.Printf("%v:%T", str1, str1)
	str2 := "An implicityly typed string"
	fmt.Printf("%v, : %T \n", str2, str2)
	fmt.Println(strings.ToUpper(str1))
	fmt.Println(strings.Title(str1))

	lValue := "hello"
	uValue := "HELLO"

	fmt.Printf("Equals case sensitive %v \n", (lValue == uValue))
	fmt.Printf("Equals s non case sensitive %v \n", strings.EqualFold(lValue, uValue))

	fmt.Println("contains exp?", strings.Contains(str1, "exp"))
	fmt.Println("contains exp?", strings.Contains(str2, "exp"))
}
