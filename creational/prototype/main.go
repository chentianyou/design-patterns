////////////////////////////////////////////////////////////////////////////////
// Copyright 2016, Oushu Inc.
// All rights reserved.
//
// Author    : chentianyou
// Create At : 2021-04-20 18:08
////////////////////////////////////////////////////////////////////////////////

package main

import "fmt"

type INode interface {
	Print(indentation string)
	Clone() INode
}

type File struct {
	Name string
}

func (f *File) Print(indentation string) {
	fmt.Println(indentation + f.Name)
}

func (f *File) Clone() INode {
	return &File{Name: f.Name + "_clone"}
}

type Folder struct {
	Children []INode
	Name     string
}

func (f *Folder) Print(indentation string) {
	fmt.Println(indentation + f.Name)
	for _, i := range f.Children {
		i.Print(indentation + indentation)
	}
}

func (f *Folder) Clone() INode {
	cloneFolder := &Folder{Name: f.Name + "_clone"}
	var tempChildren []INode
	for _, i := range f.Children {
		copy := i.Clone()
		tempChildren = append(tempChildren, copy)
	}
	cloneFolder.Children = tempChildren
	return cloneFolder
}

func main() {
	file1 := &File{Name: "File1"}
	file2 := &File{Name: "File2"}
	file3 := &File{Name: "File3"}

	folder1 := &Folder{
		Children: []INode{file1},
		Name:      "Folder1",
	}

	folder2 := &Folder{
		Children: []INode{folder1, file2, file3},
		Name:      "Folder2",
	}
	fmt.Println("\nPrinting hierarchy for Folder2")
	folder2.Print("  ")

	cloneFolder := folder2.Clone()
	fmt.Println("\nPrinting hierarchy for clone Folder")
	cloneFolder.Print("  ")
}
