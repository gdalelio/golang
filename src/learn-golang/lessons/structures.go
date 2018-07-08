package lessons

import "fmt"

//Vertex is a structuture for points
type Vertex struct {
	x int
	y int
}

//Struc uses vertex structure
func Struc() {
	point := Vertex{}
	fmt.Println(point)
	fmt.Println("The x value is: ", point.x)
	fmt.Println("The y value is: ", point.y)
	fmt.Printf("The vertex is located at: (%d, %d)\n", point.x, point.y)
	fmt.Println("printing vertexes v1,p, v2, and v3")
	fmt.Println(v1, p, v2, v3)
	fmt.Printf("p.x = %d\n", p.x)
	fmt.Printf("p = %d\n", p)
	p2.x = 4
	fmt.Printf("p2.x = %d\n", p2.x)
	fmt.Printf("p = %d\n", p)

}

var (
	v1 = Vertex{1, 2}
	v2 = Vertex{x: 1}
	v3 = Vertex{}
	p  = &Vertex{1, 2}
	p2 = p
)
