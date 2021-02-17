//to run this go to directory of main file and run command in terminal go run *.go
package main

import (
	"fmt"
)

func main(){
	adidasFactory,_ := GetSportFactory("adidas")
	nikeFactory,_ := GetSportFactory("nike")

	adidasShoe := adidasFactory.MakeShoe()
	adidasShoe.SetLogo("adidas")
	adidasShoe.SetSize(12)

	adidasShort := adidasFactory.MakeShort()

	nikeShoe  := nikeFactory.MakeShoe()
	nikeShort := nikeFactory.MakeShort()

	PrintShoeDetails(nikeShoe)
	PrintShortDetails(nikeShort)

	PrintShoeDetails(adidasShoe)
	PrintShortDetails(adidasShort)
}

func PrintShoeDetails(s Shoe){
	fmt.Println("logo of shoe=",s.GetLogo())
	fmt.Println("size of shoe=",s.GetSize())
}

func  PrintShortDetails(s Short){
	fmt.Println("logo of short=",s.GetLogo())
	fmt.Println("size of short=",s.GetSize())
}
