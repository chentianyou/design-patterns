////////////////////////////////////////////////////////////////////////////////
// Copyright 2016, Oushu Inc.
// All rights reserved.
//
// Author    : chentianyou
// Create At : 2021-04-20 10:56
////////////////////////////////////////////////////////////////////////////////

package singleton

import "sync"

type singleton struct {}

var instance *singleton
var once sync.Once

func GetInstance() *singleton {
	once.Do(func() {
		instance = &singleton{}
	})
	return instance
}
