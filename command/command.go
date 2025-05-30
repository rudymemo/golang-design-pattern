package command

type Command interface {
	Execute() string
	Undo() string
}

type Light struct {
	isOn bool
}

func NewLight() *Light {
	return &Light{isOn: false}
}

func (l *Light) TurnOn() string {
	l.isOn = true
	return "Light is ON"
}

func (l *Light) TurnOff() string {
	l.isOn = false
	return "Light is OFF"
}

func (l *Light) IsOn() bool {
	return l.isOn
}

type LightOnCommand struct {
	light *Light
}

func NewLightOnCommand(light *Light) *LightOnCommand {
	return &LightOnCommand{light: light}
}

func (loc *LightOnCommand) Execute() string {
	return loc.light.TurnOn()
}

func (loc *LightOnCommand) Undo() string {
	return loc.light.TurnOff()
}

type LightOffCommand struct {
	light *Light
}

func NewLightOffCommand(light *Light) *LightOffCommand {
	return &LightOffCommand{light: light}
}

func (loff *LightOffCommand) Execute() string {
	return loff.light.TurnOff()
}

func (loff *LightOffCommand) Undo() string {
	return loff.light.TurnOn()
}

type RemoteControl struct {
	command Command
	history []Command
}

func NewRemoteControl() *RemoteControl {
	return &RemoteControl{
		history: make([]Command, 0),
	}
}

func (rc *RemoteControl) SetCommand(command Command) {
	rc.command = command
}

func (rc *RemoteControl) PressButton() string {
	if rc.command == nil {
		return "No command set"
	}
	
	result := rc.command.Execute()
	rc.history = append(rc.history, rc.command)
	return result
}

func (rc *RemoteControl) PressUndo() string {
	if len(rc.history) == 0 {
		return "No command to undo"
	}
	
	lastCommand := rc.history[len(rc.history)-1]
	rc.history = rc.history[:len(rc.history)-1]
	return lastCommand.Undo()
}

func (rc *RemoteControl) GetHistorySize() int {
	return len(rc.history)
}