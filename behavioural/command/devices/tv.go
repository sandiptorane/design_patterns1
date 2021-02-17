package devices

import "fmt"

type TV struct {
	isRunning bool
}

func (tv *TV) On(){
	tv.isRunning = true
	fmt.Println("Turning on tv")
}

func (tv *TV) Off(){
	tv.isRunning = false
	fmt.Println("Turning off tv")
}