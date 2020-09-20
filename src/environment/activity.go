package environment

var AvailableActivity = make(map[string]string)

func ActivityToStringArray() []string {
	tab := make([]string, 0)
	for act := range AvailableActivity {
		tab = append(tab, act)
	}
	return tab
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
	env.autoAddCalendar = append(env.autoAddCalendar, activities...)
}

func (env *Environment) AddAutoRegisterActivity(activities ...string) {
	env.autoRegister = append(env.autoRegister, activities...)
}

func (env Environment) IsAutoRegisteredActivity(activityName string) bool {
	return isElementPresent(env.autoRegister, activityName)
}

func (env Environment) IsAutoCalendarRegisteredActivity(activityName string) bool {
	return isElementPresent(env.autoAddCalendar, activityName)
}
