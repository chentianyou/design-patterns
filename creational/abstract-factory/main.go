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

type Color interface {
	Fill()
}

type Red struct {
	Code string
}

func (r *Red) Fill() {
	fmt.Println("Inside Red::fill() method.")
}

type Green struct {
	Code string
}

func (r *Green) Fill() {
	fmt.Println("Inside Green::fill() method.")
}

type Blue struct {
	Code string
}

func (r *Blue) Fill() {
	fmt.Println("Inside Blue::fill() method.")
}

const (
	ColorRed   = "RED"
	ColorGreen = "GREEN"
	ColorBlue  = "BLUE"
)

type AbstractFactory interface {
	GetColor(color string) Color
	GetShape(shape string) Shape
}

type ShapeFactory struct{}

func (f ShapeFactory) GetColor(color string) Color {
	return nil
}

func (f ShapeFactory) GetShape(shape string) Shape {
	switch shape {
	case ShapeRectangle:
		return &Rectangle{}
	case ShapeSquare:
		return &Square{}
	case ShapeCircle:
		return &Circle{}
	}
	return nil
}

type ColorFactory struct{}

func (f ColorFactory) GetColor(color string) Color {
	switch color {
	case ColorRed:
		return &Red{}
	case ColorGreen:
		return &Green{}
	case ColorBlue:
		return &Blue{}
	}
	return nil
}

func (f ColorFactory) GetShape(shape string) Shape {
	return nil
}

const (
	FactoryColor = "COLOR"
	FactoryShape = "SHAPE"
)
type FactoryProducer struct {}

func (f FactoryProducer) GetFactory(choice string) AbstractFactory {
	switch choice {
	case FactoryColor:
		return &ColorFactory{}
	case FactoryShape:
		return &ShapeFactory{}
	}
	return nil
}

func main() {
	factoryProducer := &FactoryProducer{}
	shapeFactory := factoryProducer.GetFactory(FactoryShape)
	shape1 := shapeFactory.GetShape(ShapeRectangle)
	shape1.Draw()
	shape2 := shapeFactory.GetShape(ShapeSquare)
	shape2.Draw()
	shape3 := shapeFactory.GetShape(ShapeCircle)
	shape3.Draw()

	colorFactory := factoryProducer.GetFactory(FactoryColor)
	color1 := colorFactory.GetColor(ColorRed)
	color1.Fill()

	color2 := colorFactory.GetColor(ColorGreen)
	color2.Fill()

	color3 := colorFactory.GetColor(ColorBlue)
	color3.Fill()
}