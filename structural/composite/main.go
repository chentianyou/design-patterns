////////////////////////////////////////////////////////////////////////////////
// Copyright 2016, Oushu Inc.
// All rights reserved.
//
// Author    : chentianyou
// Create At : 2021-04-23 10:25
////////////////////////////////////////////////////////////////////////////////

package main

import "fmt"

type Employee struct {
	name         string
	dept         string
	salary       int
	subordinates []*Employee
}

func NewEmployee(name, dept string, sal int) *Employee {
	return &Employee{
		name:   name,
		dept:   dept,
		salary: sal,
	}
}

func (e *Employee) Add(em *Employee) {
	e.subordinates = append(e.subordinates, em)
}

func (e *Employee) Remove(em *Employee) {
	idx := -1
	for i, s := range e.subordinates {
		if s == em {
			idx = i
			break
		}
	}
	e.subordinates = append(e.subordinates[:idx], e.subordinates[idx+1:]...)
}

func (e *Employee) GetSubordinates() []*Employee {
	return e.subordinates
}

func (e *Employee) String() string {
	return fmt.Sprintf("Employee :[ Name : %s, dept : %s, salary : %d]", e.name, e.dept, e.salary)
}

func main() {
	CEO := NewEmployee("John", "CEO", 30000)

	headSales := NewEmployee("Robert", "Head Sales", 20000)

	headMarketing := NewEmployee("Michel", "Head Marketing", 20000)

	clerk1 := NewEmployee("Laura", "Marketing", 10000)
	clerk2 := NewEmployee("Bob", "Marketing", 10000)

	salesExecutive1 := NewEmployee("Richard", "Sales", 10000)
	salesExecutive2 := NewEmployee("Rob", "Sales", 10000)

	CEO.Add(headSales)
	CEO.Add(headMarketing)

	headSales.Add(salesExecutive1)
	headSales.Add(salesExecutive2)

	headMarketing.Add(clerk1)
	headMarketing.Add(clerk2)

	//打印该组织的所有员工
	fmt.Println(CEO)
	for _, headEmployee := range CEO.GetSubordinates() {
		fmt.Println(headEmployee)
		for _, employee := range headEmployee.GetSubordinates() {
			fmt.Println(employee)
		}
	}
}
