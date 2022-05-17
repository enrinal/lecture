package light

type LightOffState struct {
	LightSwitch *LightSwitch
}

func (l *LightOffState) LightPress() {
	l.LightSwitch.ChangeState(&LightOnState{
		LightSwitch: l.LightSwitch,
	})
}

func (l *LightOffState) CanLightTurnOn() bool {
	return true
}

func (l *LightOffState) GetLightCondition() string {
	return "light is off"
}
