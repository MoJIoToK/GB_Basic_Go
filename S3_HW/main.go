package main

import (
	"fmt"
	"math"
)

type Shape struct {
	Name string
}

type Rectangle struct {
	Shape
	Length float64
	Width  float64
}

type Circle struct {
	Shape
	Radius float64
}

func (c Circle) Area() float64 {
	return math.Pi * c.Radius * c.Radius
}

func (r Rectangle) Area() float64 {
	return r.Width * r.Length
}

func (s Shape) GetName() string {
	return s.Name
}

func (s Shape) Area() float64 {
	return 0.0
}

func main() {
	circle := Circle{Shape{"Круг"}, 5}
	rectangle := Rectangle{Shape{"Прямоугольник"}, 4, 6}

	fmt.Printf("%s: Площадь = %.2f\n", circle.GetName(), circle.Area())
	fmt.Printf("%s: Площадь = %.2f\n", rectangle.GetName(), rectangle.Area())
}
