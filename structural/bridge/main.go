////////////////////////////////////////////////////////////////////////////////
// Copyright 2016, Oushu Inc.
// All rights reserved.
//
// Author    : chentianyou
// Create At : 2021-04-22 10:45
////////////////////////////////////////////////////////////////////////////////

package main

import "fmt"

// Step 1
// 创建桥接实现接口。
type DrawAPI interface {
	DrawCircle(radius, x, y int)
}

// Step 2
// 创建实现了 DrawAPI 接口的实体桥接实现类。
type RedCircle struct{}

func (c *RedCircle) DrawCircle(radius, x, y int) {
	fmt.Printf("Drawing Circle[ color: red, radius: %d, x: %d, %d]\n", radius, x, y)
}

type GreenCircle struct{}

func (c *GreenCircle) DrawCircle(radius, x, y int) {
	fmt.Printf("Drawing Circle[ color: green, radius: %d, x: %d, %d]\n", radius, x, y)
}

// Step 3
// 使用 DrawAPI 接口创建抽象类 Shape。
type Shape struct {
	drawAPI DrawAPI
}

func NewShape(api DrawAPI) Shape {
	return Shape{drawAPI: api}
}

func (s *Shape) Draw() {
	panic("implement Draw")
}

// Step 4
// 创建实现了 Shape 抽象类的实体类。
type Circle struct {
	Shape
	X      int
	Y      int
	Radius int
}

func NewCircle(x, y, radius int, api DrawAPI) *Circle {
	c := &Circle{
		X:      x,
		Y:      y,
		Radius: radius,
	}
	c.Shape = NewShape(api)
	return c
}

func (c *Circle) Draw() {
	c.drawAPI.DrawCircle(c.Radius, c.X, c.Y)
}

func main() {
	redCircle := NewCircle(100, 100, 10, new(RedCircle))
	greenCircle := NewCircle(100, 100, 10, new(GreenCircle))

	redCircle.Draw()
	greenCircle.Draw()
}
