package environment

const (
	ActivityConference  = "Conference"
	ActivityDefense     = "Defense"
	ActivityDelivery    = "Delivery"
	ActivityPitch       = "Pitch"
	ActivityProject     = "Project"
	ActivityRush        = "Rush"
	ActivityTP          = "TP"
	ActivityKickOff     = "Kick-off"
	ActivityProjectTime = "Project time"
)

func (env *Environment) AddAutoRegisterActivity(activities ...string) {
	for _, activity := range activities {
		env.autoRegister = append(env.autoRegister, activity)
	}
}

func (env Environment) IsAutoRegisteredActivity(activityName string) bool {
	if len(env.autoRegister) == 0 {
		return false
	}
	for _, activity := range env.autoRegister {
		if activity == activityName {
			return true
		}
	}
	return false
}
