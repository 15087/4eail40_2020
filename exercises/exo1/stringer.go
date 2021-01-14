package main

import "fmt"

// Implement types Rectangle, Circle and Shape
type Rectangle struct {
	Width  float64
	Length float64
}

type Circle struct {
	Radius float64
}
type Shape interface {
	Area() float64
   	fmt.Stringer
}
func (r Rectangle) Area() float64{
    return r.Width  * r.Length
}
func (c Circle) Area() float64{
    return 3.14*math.Pow(c.r,2)
}
func (r Rectangle) String() string {
	return fmt.Sprintf("Rectangle of width %d and length %d", r.Width, r.Length)
}

func (c Circle) String() string {
	return fmt.Sprintf("Circle of radius %d", c.Radius)
}



func main() {
	r := &Rectangle{2, 3}
	c := &Circle{5}

	shapes := []Shape{r, c}

	for _, s := range shapes {
		fmt.Println(s)
        	fmt.Println("Area is:",s.Area())
		
	}
}
