package departments

type ComputerScience struct {
	NumOfProfessors int
}

func (c *ComputerScience)GetNumberOfProfessors() int{
	return c.NumOfProfessors
}

func (c *ComputerScience)GetName() string{
	return "computerscience"
}