////////////////////////////////////////////////////////////////////////////////
// Copyright 2016, Oushu Inc.
// All rights reserved.
//
// Author    : chentianyou
// Create At : 2021-04-20 11:17
////////////////////////////////////////////////////////////////////////////////

package main

import "fmt"

type Item interface {
	Name() string
	Packing() Packing
	Price() float32
}

type Packing interface {
	Pack() string
}

type Wrapper struct{}

func (w *Wrapper) Pack() string {
	return "Wrapper"
}

type Bottle struct{}

func (b *Bottle) Pack() string {
	return "Bottle"
}

type Burger struct{}

func (b *Burger) Packing() Packing {
	return &Wrapper{}
}

func (b *Burger) Name() string {
	panic("implements this function")
}

func (b *Burger) Price() float32 {
	panic("implements this function")
}

type ColdDrink struct{}

func (b *ColdDrink) Packing() Packing {
	return &Bottle{}
}

func (b *ColdDrink) Name() string {
	panic("implements this function")
}

func (b *ColdDrink) Price() float32 {
	panic("implements this function")
}

type VegBurger struct {
	Burger
}

func (v *VegBurger) Price() float32 {
	return 25.0
}

func (v *VegBurger) Name() string {
	return "Veg Burger"
}

type ChickenBurger struct {
	Burger
}

func (v *ChickenBurger) Price() float32 {
	return 50.5
}

func (v *ChickenBurger) Name() string {
	return "Chicken Burger"
}

type Coke struct {
	ColdDrink
}

func (v *Coke) Price() float32 {
	return 30.0
}

func (v *Coke) Name() string {
	return "Coke"
}

type Pepsi struct {
	ColdDrink
}

func (v *Pepsi) Price() float32 {
	return 35.0
}

func (v *Pepsi) Name() string {
	return "Pepsi"
}

type Meal struct {
	Items []Item
}

func (m *Meal) AddItem(item Item) {
	m.Items = append(m.Items, item)
}

func (m *Meal) GetCost() float32 {
	cost := float32(0.0)
	for _, item := range m.Items {
		cost += item.Price()
	}
	return cost
}

func (m *Meal) ShowItems() {
	for _, item := range m.Items {
		fmt.Print("Item : " + item.Name())
		fmt.Print(", Packing : " + item.Packing().Pack())
		fmt.Println(", Price : ", item.Price())
	}
}

type MealBuilder struct{}

func (m *MealBuilder) PrepareVegMeal() *Meal {
	meal := new(Meal)
	meal.AddItem(&VegBurger{})
	meal.AddItem(&Coke{})
	return meal
}

func (m *MealBuilder) PrepareNonVegMeal() *Meal {
	meal := new(Meal)
	meal.AddItem(&ChickenBurger{})
	meal.AddItem(&Pepsi{})
	return meal
}

func main() {
	mealBuilder := &MealBuilder{}
	vegMeal := mealBuilder.PrepareVegMeal()
	fmt.Println("Veg Meal")
	vegMeal.ShowItems()
	fmt.Println("Total Cost: ", vegMeal.GetCost())

	nonVegMeal := mealBuilder.PrepareNonVegMeal()
	fmt.Println("\n\nNon-Veg Meal")
	nonVegMeal.ShowItems()
	fmt.Println("Total Cost: ", nonVegMeal.GetCost())
}
