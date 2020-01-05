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

func updateTable(table []string, entries ...string) []string {
	for _, entry := range entries {
		table = append(table, entry)
	}
	return table
}

func isElementPresent(reference []string, element string) bool {
	if len(reference) == 0 {
		return false
	}
	for _, activity := range reference {
		if activity == element {
			return true
		}
	}
	return false
}

func (env *Environment) AddAutoRegisterCalendarActivity(activities ...string) {
	env.autoAddCalendar = updateTable(env.autoAddCalendar, activities...)
}

func (env *Environment) AddAutoRegisterActivity(activities ...string) {
	env.autoRegister = updateTable(env.autoRegister, activities...)
}

func (env Environment) IsAutoRegisteredActivity(activityName string) bool {
	return isElementPresent(env.autoRegister, activityName)
}

func (env Environment) IsAutoCalendarRegisteredActivity(activityName string) bool {
	return isElementPresent(env.autoAddCalendar, activityName)
}
