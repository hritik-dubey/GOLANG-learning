package main

import (
    "fmt"
)

// Defining the Shape interface
type Shape interface {
    Area() float64
    Perimeter() float64
}

// Rectangle struct
type Rectangle struct {
    Width  float64
    Height float64
}

// Implementing the Shape interface for Rectangle
func (r Rectangle) Area() float64 {
    return r.Width * r.Height
}

func (r Rectangle) Perimeter() float64 {
    return 2 * (r.Width + r.Height)
}

// Circle struct
type Circle struct {
    Radius float64
}

// Implementing the Shape interface for Circle
func (c Circle) Area() float64 {
    return 3.14 * c.Radius * c.Radius
}

func (c Circle) Perimeter() float64 {
    return 2 * 3.14 * c.Radius
}

// Function to print shape information
// func PrintShapeInfo(s Shape) {
//     fmt.Printf("Area: %f\n", s.Area())
//     fmt.Printf("Perimeter: %f\n", s.Perimeter())
// }

func main() {
    rect := Rectangle{Width: 10, Height: 5}
    // circle := Circle{Radius: 7}

    fmt.Println("Rectangle:")
    // PrintShapeInfo(rect)
	area:=rect.Area()
	fmt.Println(area)

    // fmt.Println("\nCircle:")
    // PrintShapeInfo(circle)
}