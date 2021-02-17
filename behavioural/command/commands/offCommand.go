package commands

import "DesignPatterns/behavioural/command/devices"

type OffCommand struct {
	Device devices.Device
}

func (c *OffCommand) Execute() {
	c.Device.Off()
}
