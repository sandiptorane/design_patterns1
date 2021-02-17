package departments

type Mechanical struct {
	NumOfProfessors int
}

func (c *Mechanical)GetNumberOfProfessors() int{
	return c.NumOfProfessors
}

func (c *Mechanical)GetName() string{
	return "mechanical"
}
