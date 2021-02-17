package Factory

import "fmt"

func GetGun(name string) (IGun,error){
	if name=="ak47"{
		return NewAk47(),nil
	}
	if name=="maverick"{
		return NewMaverick(),nil
	}
	return nil,fmt.Errorf("wrong gun type passed")
}
