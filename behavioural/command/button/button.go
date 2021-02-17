package button

import "DesignPatterns/behavioural/command/commands"

type Button struct {
	Command commands.Command
}

func (b *Button)Press(){
	b.Command.Execute()
}
