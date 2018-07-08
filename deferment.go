package lessons

import "fmt"

//Deferer is a deeferment function to test
//the defer capabilities of Go.
func Deferer() {

	for i := 0; i < 10; i++ {
		defer fmt.Println(i)
	}
	fmt.Println("Done")

	fmt.Printf("the back of this statement was ")
	defer fmt.Println("defered!")
}
