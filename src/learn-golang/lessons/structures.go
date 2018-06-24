package lessons

import "fmt"

//Vertex is a structuture for points
type Vertex struct {
	x int
	y int
}

//Struc uses vertex structure
func Struc(x int, y int) {
	point := Vertex{x, y}
	fmt.Println(point)
	fmt.Println("The x value is: ", point.x)
	fmt.Println("The y value is: ", point.y)
	fmt.Printf("The vertex is located at: (%d, %d)\n", point.x, point.y)
}
