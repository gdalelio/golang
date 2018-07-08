package lessons

import "fmt"

//Slicing - more fun with slices
func Slicing() {
	//make creates a zeroed array and returns a slice that
	//references the array
	a := make([]int, 5) //len(a) = 5 zero capacity
	fmt.Printf("Slice a len: %d\n", len(a))
	fmt.Printf("Slice a cap: %d\n", cap(a))
	//to specify capcaity, a thir parameter is needed
	b := make([]int, 0, 5) //len(b) = 0, cap(b) = 5
	fmt.Printf("Slice b len: %d, cap: %d\n", len(b), cap(b))
	b = b[:cap(b)] //len(b) = 5, cap(b) = 5
	fmt.Printf("Slice b len: %d, cap: %d\n", len(b), cap(b))
	b = b[1:]
	fmt.Printf("Slice b len: %d, cap: %d\n", len(b), cap(b))

}
