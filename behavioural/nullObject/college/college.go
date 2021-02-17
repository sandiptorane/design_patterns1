package college

import dept "DesignPatterns/behavioural/nullObject/departments"

type College struct {
	Departments  []dept.Department
}


func (c *College) AddDepartment(departmentName string, numOfProfessors int) {
	if departmentName == "computerscience" {
		computerScienceDepartment := &dept.ComputerScience{NumOfProfessors: numOfProfessors}
		c.Departments = append(c.Departments, computerScienceDepartment)
	}
	if departmentName == "mechanical" {
		mechanicalDepartment := &dept.Mechanical{NumOfProfessors: numOfProfessors}
		c.Departments = append(c.Departments, mechanicalDepartment)
	}
	return
}

func (c *College) GetDepartment(departmentName string) dept.Department {
	for _, department := range c.Departments {
		if department.GetName() == departmentName {
			return department
		}
	}
	//Return a null department if the department doesn't exits
	return &dept.NullDepartment{}
}