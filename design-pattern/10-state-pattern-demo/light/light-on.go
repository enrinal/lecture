package light

type LightOnState struct {
	LightSwitch *LightSwitch
}

func (l *LightOnState) LightPress() {
	l.LightSwitch.ChangeState(&LightOffState{
		LightSwitch: l.LightSwitch,
	})
}

func (l *LightOnState) CanLightTurnOn() bool {
	return false
}

func (l *LightOnState) GetLightCondition() string {
	return "light is on"
}
