package main

import (
	"DesignPatterns/behavioural/command/button"
	"DesignPatterns/behavioural/command/commands"
	"DesignPatterns/behavioural/command/devices"
)

func main(){
	tv := &devices.TV{}

   onCommand := &commands.OnCommand{
   	  Device : tv,
   }
   offCommand := &commands.OffCommand{
   	Device: tv,
   }

   onButton := &button.Button{
   	Command: onCommand,
   }
   onButton.Press()

   offButton := &button.Button{
   	Command: offCommand,
   }
   offButton.Press()
}
