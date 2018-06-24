package exercises

import "fmt"

//Sqrt - determines the square root of x by approximation
func Sqrt(x float64) float64 {

	var z = 1.0
	var lastZ float64

	fmt.Printf("Starting vals =>lastZ = %f z = %f\n", lastZ, z)

	for i := 1; i < 10; i++ {
		z -= (z*z - x) / (2 * z)
		fmt.Printf("lastZ: %g z: %g loop #: %d\n", lastZ, z, i)
		if lastZ == z {
			return z
		} else {

			lastZ = z
		}

	}
	return z
}
