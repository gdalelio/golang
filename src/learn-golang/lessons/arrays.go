package lessons

import "fmt"

//Arrays function for learning array features
func Arrays() {
	var a [2]string
	a[0] = "Hello"
	a[1] = "Geoff"
	fmt.Println("--------Fun with Arrays-------")
	fmt.Println(a[0], a[1])
	fmt.Println(a)
	primes := [6]int{2, 3, 5, 7, 11, 13}
	fmt.Println(primes)
	fmt.Printf("primes array can hold %d members.\n", len(primes))
	fmt.Println("--------Fun with Slices-------")
	fmt.Println("pulling part of the primes array (elements primes[2] & primes[3])")
	//[low bound : high bound] - starts at low and excludes upper bound element
	var slice1 []int = primes[2:4] //elements 2 -3
	fmt.Println(slice1)
	r := [3]bool{true, true, false}
	fmt.Println(r)
	s := []int{1, 2, 3, 4, 5, 6, 7}
	fmt.Println(s)
	s1 := s[:3]
	fmt.Printf("slice of s ->s1[ : 3] ==>%d\n", s1)
	fmt.Printf("Slice length: %d, Slice capacity: %d\n", len(s1), cap(s1))
	//resizing s1 to 5 elements
	s1 = s[0:6]
	fmt.Printf("slice of s ->s1[ : 6] ==>%d\n", s1)
	fmt.Printf("Slice length: %d, Slice capacity: %d\n", len(s1), cap(s1))

	//the zero value of a slice is nil
	//a nil slice has a length and capacity of 0 and no underlying Arrays
	var slice2 []int
	fmt.Printf("slice2: %s -> length of slice2 is %d, capacity of slice2 is %d\n", slice2, len(slice2), cap(slice2))
	if slice2 == nil {
		fmt.Println("Slice2 is nil!")
	}
}
