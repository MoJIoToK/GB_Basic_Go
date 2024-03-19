package main

import (
	"fmt"
	"math"
)

type Name string

type Shape interface {
	GetArea() float64
	GetName() string
}

type Rectangle struct {
	name   string
	Length float64
	Width  float64
}

type Circle struct {
	name   string
	Radius float64
}

func (c Circle) GetArea() float64 {
	return math.Pi * c.Radius * c.Radius
}

func (c Circle) GetName() string {
	c.name = "Круг"
	return c.name
}

func (r Rectangle) GetArea() float64 {
	return r.Width * r.Length
}

func (r Rectangle) GetName() string {
	r.name = "Прямоугольник"
	return r.name
}

func main() {
	circle := Circle{Radius: 5}
	//circle := Circle{Shape{"Круг"}, 5}
	//rectangle := Rectangle{Shape{"Прямоугольник"}, 4, 6}

	fmt.Printf("%s: Площадь = %.2f\n", circle.GetName(), circle.GetArea())
	//fmt.Printf("%s: Площадь = %.2f\n", rectangle.GetName(), rectangle.GetArea())
}
