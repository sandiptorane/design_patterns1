package main

import (
	"fmt"
	"sync"
)

var once sync.Once

type Single struct{

}

var SingleInstance *Single

func GetInstance() *Single{
	if singleInstance==nil{
		once.Do(func(){
			fmt.Println("Creating single instance")
			SingleInstance = &Single{}
		})
	}else{
		fmt.Println("sinlgeInstance already created")
	}
	return SingleInstance
}

func main(){
	for i:=0;i<30;i++{
		go GetInstance()
	}
	fmt.Scanln()
}
