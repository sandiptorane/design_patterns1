package departments

import (
	"DesignPatterns/behavioural/chainOfResponsibility/patient"
	"fmt"
)

type Cashier struct {
	next Department
}

func (c *Cashier) Execute(p *patient.Patient) {
	if p.PaymentDone {
		fmt.Println("Payment Done")
	}
	fmt.Println("Cashier getting money from patient patient")
}

func (c *Cashier) Next(next Department) {
	c.next = next
}

