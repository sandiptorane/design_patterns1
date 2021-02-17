package main

type iglooBuilder struct {
	windowType string
	doorType   string
	floor      int
}
func NewIglooBuilder() *normalBuilder{
	return &normalBuilder{}
}

func (n *iglooBuilder)SetWindowType(){
	n.windowType= "snow window"
}


func (n *iglooBuilder)SetDoorType(){
	n.doorType = "snow door"
}

func (n *iglooBuilder)SetNumFloor(){
	n.floor = 1
}

func (n *iglooBuilder)GetHouse() house{
	return house{
		doorType: n.doorType,
		windowType: n.windowType,
		floor: n.floor,
	}
}

