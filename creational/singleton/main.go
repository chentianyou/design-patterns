////////////////////////////////////////////////////////////////////////////////
// Copyright 2016, Oushu Inc.
// All rights reserved.
//
// Author    : chentianyou
// Create At : 2021-04-20 10:55
////////////////////////////////////////////////////////////////////////////////

package main

import (
	"fmt"
	"github.com/chentianyou/design-patterns/creational/singleton/singleton"
)

func main() {
	instance1 := singleton.GetInstance()
	fmt.Printf("instance1 address is %p\n", instance1)

	instance2 := singleton.GetInstance()
	fmt.Printf("instance2 address is %p\n", instance2)
}