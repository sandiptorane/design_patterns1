package main

import "fmt"

func main(){
	normalBuilder := GetBuilder("normal")
	iglooBuilder  := GetBuilder("igloo")

	director := NewDirector(normalBuilder)
	noramlHouse := director.BuildHouse()
	fmt.Println("normal house door type :",noramlHouse.doorType)
	fmt.Println("normal house window type:",noramlHouse.windowType)
	fmt.Println("normal house has no of floors :",noramlHouse.floor)

	director.SetBuilder(iglooBuilder)
	iglooHouse := director.BuildHouse()
	fmt.Println("igloo house door type :",iglooHouse.doorType)
	fmt.Println("igloo house window type :",iglooHouse.windowType)
	fmt.Println("igloo house has floors :",iglooHouse.floor)

}
