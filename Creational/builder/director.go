package main

type director struct {
	builder iBuilder
}

func NewDirector(b iBuilder) *director{
	return &director{
		builder : b,
	}
}

func (d *director)SetBuilder(b iBuilder){
	d.builder =b
}

func (d *director)BuildHouse() house{
	d.builder.SetDoorType()
	d.builder.SetWindowType()
	d.builder.SetNumFloor()
	return d.builder.GetHouse()
}