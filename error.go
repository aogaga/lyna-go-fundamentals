package main

import(
	"fmt"
	"github.com/pkg/errors"
	"os"
)

func main()  {

	f, err := os.Open("filenam.txt")

	if(err == nil){
		fmt.Println(f)
	}else{
		fmt.Println(err.Error())
	}

	myError := errors.New(" my error string")
	fmt.Println(myError)
}