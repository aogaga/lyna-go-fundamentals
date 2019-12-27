package main

import(
	"fmt"
)


func main(){
	str1 := "The quick red fox"
	str2 := "jumped over"
	str3 := "the lazy brown bag"
	aNumber := 42
	isTrue := true

	stringLength, err := fmt.Println(str1, str2, str3)
	if err == nil{
		fmt.Println("string lenght", stringLength);
	}

	fmt.Printf("Value of a number %v\n", aNumber)
	fmt.Printf("The value of isTrue is %v\n", isTrue)
	fmt.Printf("value os aNumber as float: %.2f\n" , float64(aNumber))

	fmt.Printf("Data types: %T, %T, %T, %T,  and %T", str1, str2, str3, aNumber, isTrue )

	myString := fmt.Sprintf("Data Var: %T, %T, %T, %T,  and %T", str1, str2, str3, aNumber, isTrue )
	fmt.Printf(myString)

}