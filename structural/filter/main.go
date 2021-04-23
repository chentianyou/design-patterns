////////////////////////////////////////////////////////////////////////////////
// Copyright 2016, Oushu Inc.
// All rights reserved.
//
// Author    : chentianyou
// Create At : 2021-04-23 09:45
////////////////////////////////////////////////////////////////////////////////

package main

import (
	"fmt"
	"strings"
)

// Step 1
// 创建一个类，在该类上应用标准。
type Person struct {
	name          string
	gender        string
	maritalStatus string
}

func NewPerson(name, gender, maritalStatus string) *Person {
	return &Person{
		name:          name,
		gender:        gender,
		maritalStatus: maritalStatus,
	}
}

func (p *Person) GetName() string {
	return p.name
}

func (p *Person) GetGender() string {
	return p.gender
}

func (p *Person) GetMaritalStatus() string {
	return p.maritalStatus
}

// Step2
// 为标准（Criteria）创建一个接口。
type Criteria interface {
	MeetCriteria(persons []*Person) []*Person
}

// Step3
// 创建实现了 Criteria 接口的实体类。
type CriteriaMale struct{}

func (c *CriteriaMale) MeetCriteria(persons []*Person) []*Person {
	var malePersons []*Person
	for _, person := range persons {
		if strings.ToUpper(person.GetGender()) == "MALE" {
			malePersons = append(malePersons, person)
		}
	}
	return malePersons
}

type CriteriaFemale struct{}

func (c *CriteriaFemale) MeetCriteria(persons []*Person) []*Person {
	var femalePersons []*Person
	for _, person := range persons {
		if strings.ToUpper(person.GetGender()) == "FEMALE" {
			femalePersons = append(femalePersons, person)
		}
	}
	return femalePersons
}

type CriteriaSingle struct{}

func (c *CriteriaSingle) MeetCriteria(persons []*Person) []*Person {
	var singlePersons []*Person
	for _, person := range persons {
		if strings.ToUpper(person.GetMaritalStatus()) == "SINGLE" {
			singlePersons = append(singlePersons, person)
		}
	}
	return singlePersons
}

type AndCriteria struct {
	criteria      Criteria
	otherCriteria Criteria
}

func NewAndCriteria(criteria, otherCriteria Criteria) *AndCriteria {
	return &AndCriteria{
		criteria:      criteria,
		otherCriteria: otherCriteria,
	}
}

func (c *AndCriteria) MeetCriteria(persons []*Person) []*Person {
	firstCriteriaPersons := c.criteria.MeetCriteria(persons)
	return c.otherCriteria.MeetCriteria(firstCriteriaPersons)
}

type OrCriteria struct {
	criteria      Criteria
	otherCriteria Criteria
}

func NewOrCriteria(criteria, otherCriteria Criteria) *OrCriteria {
	return &OrCriteria{
		criteria:      criteria,
		otherCriteria: otherCriteria,
	}
}

func (c *OrCriteria) MeetCriteria(persons []*Person) []*Person {
	firstCriteriaItems := c.criteria.MeetCriteria(persons)
	otherCriteriaItems := c.otherCriteria.MeetCriteria(persons)
	for _, p1 := range otherCriteriaItems {
		for _, p2 := range firstCriteriaItems {
			if p1 == p2 {
				firstCriteriaItems = append(firstCriteriaItems, p1)
				break
			}
		}
	}
	return firstCriteriaItems
}

func main() {
	var persons []*Person
	persons = append(persons, NewPerson("Robert", "Male", "Single"))
	persons = append(persons, NewPerson("John", "Male", "Married"))
	persons = append(persons, NewPerson("Laura", "Female", "Married"))
	persons = append(persons, NewPerson("Diana", "Female", "Single"))
	persons = append(persons, NewPerson("Mike", "Male", "Single"))
	persons = append(persons, NewPerson("Bobby", "Male", "Single"))

	male := new(CriteriaMale)
	female := new(CriteriaFemale)
	single := new(CriteriaSingle)
	singleMale := NewAndCriteria(single, male)
	singleOrFemale := NewOrCriteria(single, female)

	fmt.Println("Males: ")
	printPersons(male.MeetCriteria(persons))

	fmt.Println("\nFemales: ")
	printPersons(female.MeetCriteria(persons))

	fmt.Println("\nSingle Males: ")
	printPersons(singleMale.MeetCriteria(persons))

	fmt.Println("\nSingle Or Females: ")
	printPersons(singleOrFemale.MeetCriteria(persons))
}

func printPersons(persons []*Person) {
	for _, person := range persons {
		fmt.Printf("Person : [ Name : %s, Gender : %s, Marital Status : %s ]\n",
			person.GetName(), person.GetGender(), person.GetMaritalStatus())
	}
}
