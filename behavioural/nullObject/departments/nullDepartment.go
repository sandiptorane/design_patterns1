package departments

type NullDepartment struct {
	NumOfProfessors int
}

func (c *NullDepartment) GetNumberOfProfessors() int {
	return 0
}

func (c *NullDepartment) GetName() string {
	return "nullDepartment"
}
