package lessons

import "fmt"

func Pointers() {
	i, j := 42, 2701

	p := &i         //point to i
	fmt.Println(*p) //read i through the Pointer

}
