package main

import(
	"fmt"
)

type Animal interface {
	Speak() string
}

type Puppy struct {

}

func (d Puppy) Speak() string  {
	return "woot !"
}

func main()  {
	poodle := Animal(Puppy{})

	fmt.Println(poodle)
}