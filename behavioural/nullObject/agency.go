package main

import (
	"DesignPatterns/behavioural/nullObject/college"
	"fmt"
)

func main() {
	college1 := createCollege1()
	college2 := createCollege2()
	totalProfessors := 0
	departmentArray := []string{"computerscience", "mechanical", "civil", "electronics"}

	for _, deparmentName := range departmentArray {
		d := college1.GetDepartment(deparmentName)
		totalProfessors += d.GetNumberOfProfessors()
	}

	fmt.Printf("Total number of professors in college1 is %d\n", totalProfessors)

	//Reset the professor count
	totalProfessors = 0
	for _, deparmentName := range departmentArray {
		d := college2.GetDepartment(deparmentName)
		totalProfessors += d.GetNumberOfProfessors()
	}
	fmt.Printf("Total number of professors in college2 is %d\n", totalProfessors)
}

func createCollege1() *college.College {
	college := &college.College{}
	college.AddDepartment("computerscience", 4)
	college.AddDepartment("mechanical", 3)
	return college
}

func createCollege2() *college.College {
	college := &college.College{}
	college.AddDepartment("computerscience", 2)
	return college
}
