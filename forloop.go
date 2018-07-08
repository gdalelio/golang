package lessons

import "fmt"

//For_loop  function called
func Forloop() {
	fmt.Println("in for For Loop lesson")

	sum1 := 0
	for i := 0; i < 10; i++ {
		sum1 += i
	}

	sum := 1
	for sum < 1000 {
		sum += sum
	}
	//for "while" Loop
	sum2 := 1
	for sum2 < 1000 {
		sum2 += sum2
	}

	fmt.Printf("sum = %d, sum1 = %d, sum2 = %d\n", sum, sum1, sum2)
}
