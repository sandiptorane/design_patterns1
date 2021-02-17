package main

import (
	"DesignPatterns/behavioural/mementoPattern/components"
	"fmt"
)

func main() {
   caretaker := &components.Caretaker{
   	MementoArray: make([]*components.Memento,0),
   }

   originator := &components.Originator{
   	State : "A",
   }

   fmt.Printf("Originator Current State: %s\n", originator.GetState())
   caretaker.AddMemento(originator.CreateMemento())

   originator.SetState("B")
   fmt.Printf("Originator Current State: %s\n", originator.GetState())
   caretaker.AddMemento(originator.CreateMemento())

   originator.SetState("C")
   fmt.Printf("Originator Current State: %s\n", originator.GetState())
   caretaker.AddMemento(originator.CreateMemento())

   originator.RestoreMemento(caretaker.GetMemento(2))
   fmt.Printf("Restored to State: %s\n", originator.GetState())

	originator.RestoreMemento(caretaker.GetMemento(0))
	fmt.Printf("Restored to State: %s\n", originator.GetState())

}
