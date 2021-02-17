package main

import "fmt"

type folder struct {
	childrens []inode
	name string
}

func (f *folder) print(indentation string) {
	fmt.Println(indentation + f.name + "_clone")
	for _, i := range f.childrens {
		i.print(indentation + indentation)
	}
}

func (f *folder) clone() inode {
	cloneFolder := &folder{name: f.name}
	var tempChildrens []inode
	for _, i := range f.childrens {
		copy := i.clone()           //The clone function returns a copy of the respective file or folder
		tempChildrens = append(tempChildrens, copy)
	}
	cloneFolder.childrens = tempChildrens
	return cloneFolder
}