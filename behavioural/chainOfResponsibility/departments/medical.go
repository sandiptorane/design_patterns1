package departments

import (
	"DesignPatterns/behavioural/chainOfResponsibility/patient"
	"fmt"
)

type Medical struct {
	next Department
}

func (m *Medical)Execute(p *patient.Patient){
	if p.MedicineDone{
		fmt.Println("Medicines already given to patient")
		m.next.Execute(p)
		return
	}
	fmt.Println("Medical giving the medicine to patient")
	p.MedicineDone = true
	m.next.Execute(p)
}

func (m *Medical)Next(next Department){
	m.next= next
}
