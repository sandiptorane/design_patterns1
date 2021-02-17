package main

import (
	dept "DesignPatterns/behavioural/chainOfResponsibility/departments"
	"DesignPatterns/behavioural/chainOfResponsibility/patient"
)

func main(){
	cashier := &dept.Cashier{}
	cashier.Next(nil)
	//Set next for medical department
	medical := &dept.Medical{}
	medical.Next(cashier)
	//Set next for doctor department
	doctor := &dept.Doctor{}
	doctor.Next(medical)
	//set next for reception department
	reception := &dept.Reception{}
	reception.Next(doctor)

	patients := &patient.Patient{
		Name: "sam",
	}

	//patient visiting
	reception.Execute(patients)
}

