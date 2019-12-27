package main

import(
	"fmt"
)

type Dog struct {
	Breed string 
	Weight int 
	Sound string
}

//Speak
func (d Dog) Speak(){
	fmt.Println(d.Sound)
}

func (d *Dog) SpeakThreeTimes()  {
	d.Sound = fmt.Sprintf("%v! %v! %v!\n", d.Sound, d.Sound, d.Sound)
	fmt.Print(d.Sound)

}


func main()  {
	poodle := Dog{"poodle", 37, "woot"}
	poodle.Speak()

	poodle.SpeakThreeTimes()
	poodle.SpeakThreeTimes()
}