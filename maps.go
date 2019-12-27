package main

import(
	"fmt"
	"sort"
)

func main()  {

	states := make(map[string]string)
	states["WA"] = "Washington"
	states["OR"] = "Oregon"
	states["CA"] = "California"
	fmt.Println(states)

	california := states["CA"]
	fmt.Println(california)


	delete(states, "OR")
	fmt.Println(states)

	for k, v := range states{
		fmt.Printf("%v : %v\n", k, v)
	}

	keys := make([] string, len(states))
	i := 0

	for k := range  states{
		keys[i] = k
		i++
	}

	fmt.Println(keys)


	sort.Strings(keys)

	fmt.Println("printing keys \n")

	for i := range keys {
		fmt.Println(states[keys[i]])
	}

}