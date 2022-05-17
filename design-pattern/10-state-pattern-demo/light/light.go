package light

type LightState interface {
	LightPress()
	GetLightCondition() string
	CanLightTurnOn() bool
}

type LightSwitch struct {
	State LightState
}

func NewLightSwitch() *LightSwitch {
	ls := &LightSwitch{}
	ls.ChangeState(&LightOffState{
		LightSwitch: ls,
	})
	return ls
}

func (l *LightSwitch) LightPress() {
	l.State.LightPress()
}

func (l *LightSwitch) CanLightTurnOn() bool {
	return l.State.CanLightTurnOn()
}

func (l *LightSwitch) GetLightCondition() string {
	return l.State.GetLightCondition()
}

func (l *LightSwitch) ChangeState(state LightState) {
	l.State = state
}
