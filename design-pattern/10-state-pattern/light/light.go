package light

type State interface {
	Press()
	CanTurnOnLight() bool
}

type LightSwitch struct {
	State State
}

func NewLightSwitch() *LightSwitch {
	ls := &LightSwitch{}
	ls.ChangeState(&LightOff{
		LightSwitch: ls,
	})
	return ls
}
func (l *LightSwitch) Press() {
	l.State.Press()
}

func (l *LightSwitch) CanTurnOnLight() bool {
	return l.State.CanTurnOnLight()
}

func (l *LightSwitch) ChangeState(state State) {
	l.State = state
}

type LightOn struct {
	LightSwitch *LightSwitch
}

func (l *LightOn) Press() {
	l.LightSwitch.ChangeState(&LightOff{
		LightSwitch: l.LightSwitch,
	})
}

func (l *LightOn) CanTurnOnLight() bool {
	return false
}

type LightOff struct {
	LightSwitch *LightSwitch
}

func (l *LightOff) Press() {
	l.LightSwitch.ChangeState(&LightOn{
		LightSwitch: l.LightSwitch,
	})
}

func (l *LightOff) CanTurnOnLight() bool {
	return true
}
