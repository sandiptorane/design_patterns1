package mediator

import (
	"fmt"
)

type GoodsTrain struct {
	Mediator Mediator
}

func (g *GoodsTrain) RequestArrival() {
	if g.Mediator.CanLand(g){
		fmt.Println("Goods train : Landing")
	}else{
		fmt.Println("Goods train : waiting")
	}
}

func (g *GoodsTrain) Departure(){
	fmt.Println("Goods Train: Leaving")
	g.Mediator.NotifyFree()
}

func (g *GoodsTrain) PermitArrival(){
	fmt.Println("GoodsTrain: Arrival Permitted. Landing")
}
