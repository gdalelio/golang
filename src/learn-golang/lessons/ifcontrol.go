package lessons

import (
	"fmt"
	"math"
)

//Sqrt2 exported function of SQRT
func Sqrt2(x float64) string {
	if x < 0 {
		fmt.Println("in if control")
		return Sqrt2(-x)
	}
	//	fmt.Println(Sqrt2(2), Sqrt2(4))
	return fmt.Sprint(math.Sqrt(x))
}
