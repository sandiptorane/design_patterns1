package commands

import "DesignPatterns/behavioural/command/devices"

type OnCommand struct {
	Device devices.Device
}

func (c *OnCommand) Execute(){
	c.Device.On()
}
