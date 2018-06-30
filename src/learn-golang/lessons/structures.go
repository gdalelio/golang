package lessons

import "fmt"

//Vertex is a structuture for points
type Vertex struct {
	x int
	y int
	z int
}

//Struc uses vertex structure
func Struc(x int, y int, z int) {
	point := Vertex{x, y, z}
	fmt.Println(point)
	fmt.Println("The x value is: ", point.x)
	fmt.Println("The y value is: ", point.y)
	fmt.Println("The z value is: ", point.z)
	fmt.Printf("The vertex is located at: (%d, %d, %d)\n", point.x, point.y, point.z)
	p := &point //pointer to point variable
	p.x = 1e9   //replace the value of x with 1e9
	fmt.Println(p.x)
}
