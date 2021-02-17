package main

import (
	f "DesignPatterns/Creational/factory/Factory"
	"fmt"
	"log"
)

func main() {
	ak47, err := f.GetGun("ak47")
	if err!=nil{
		log.Fatal(err)
	}
	printDetails(ak47)
	maverick,err := f.GetGun("maverick")
	if err!=nil{
		log.Fatal(err)
	}
	printDetails(maverick)

}

func printDetails(g f.IGun) {
	fmt.Printf("Gun: %s", g.GetName())
	fmt.Println()
	fmt.Printf("Power: %d", g.GetPower())
	fmt.Println()
}
