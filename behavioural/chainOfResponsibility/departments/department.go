package departments

import "DesignPatterns/behavioural/chainOfResponsibility/patient"

type Department interface {
	Execute(*patient.Patient)
	Next(Department)
}
