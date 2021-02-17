package main

import "fmt"

type SportFactory interface{
	MakeShoe() Shoe //Shoe is a interface
	MakeShort() Short   //Short is also interface
}


func GetSportFactory(brand string)(SportFactory, error){
	if brand=="nike"{
		return &Nike{},nil //here we return SportFactory for nike
	}
	if brand =="adidas"{
		return &Adidas{},nil
	}
	return nil,fmt.Errorf("wrong brand type passed")
}
