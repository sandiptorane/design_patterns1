package departments

import (
	"DesignPatterns/behavioural/chainOfResponsibility/patient"
	"fmt"
)

type Reception struct {
	next Department
}

func (r *Reception)Execute(p *patient.Patient){
       if p.RegistrationDone{
       	fmt.Println("Reception already done")
       	r.next.Execute(p)
       	return
	   }
	   fmt.Println("Reception registering patient")
		p.RegistrationDone = true
		r.next.Execute(p)
}

func (r *Reception)Next(next Department){
	r.next= next
}