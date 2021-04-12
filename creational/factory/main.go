package main

import "fmt"

// 创建一个接口
type Shape interface {
	Draw()
}

// 实现一个接口
type Rectangle struct {
	X      int
	Y      int
	Width  int
	Height int
}

func (r *Rectangle) Draw() {
	fmt.Println("Inside Rectangle::draw() method.")
}

type Square struct {
	X    int
	Y    int
	Size int
}

func (r *Square) Draw() {
	fmt.Println("Inside Square::draw() method.")
}

type Circle struct {
	X int
	Y int
	R int
}

func (r *Circle) Draw() {
	fmt.Println("Inside Circle::draw() method.")
}

const (
	ShapeRectangle = "Rectangle"
	ShapeSquare    = "Square"
	ShapeCircle    = "Circle"
)

type ShapeFactory struct{}

func (s ShapeFactory) GetShape(shapeType string) Shape {
	switch shapeType {
	case ShapeRectangle:
		return &Rectangle{}
	case ShapeSquare:
		return &Square{}
	case ShapeCircle:
		return &Circle{}
	}
	return nil
}

func main() {
	shapeFactory := ShapeFactory{}
	shape1 := shapeFactory.GetShape(ShapeRectangle)
	shape1.Draw()

	shape2 := shapeFactory.GetShape(ShapeSquare)
	shape2.Draw()

	shape3 := shapeFactory.GetShape(ShapeCircle)
	shape3.Draw()
}
