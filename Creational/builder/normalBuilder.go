package main

type normalBuilder struct {
	windowType string
	doorType   string
	floor      int
}

func NewNormalBuilder() *normalBuilder{
	return &normalBuilder{}
}

func (n *normalBuilder)SetWindowType(){
	n.windowType= "wooden window"
}


func (n *normalBuilder)SetDoorType(){
	n.doorType = "wooden door"
}

func (n *normalBuilder)SetNumFloor(){
	n.floor = 2
}

func (n *normalBuilder)GetHouse() house{
	return house{
		doorType: n.doorType,
		windowType: n.windowType,
		floor: n.floor,
	}
}
