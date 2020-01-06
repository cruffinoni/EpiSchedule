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

	/*
	   1 -> #a4bdfc
	   2 -> #7ae7bf
	   3 -> #dbadff
	   4 -> #ff887c
	   5 -> #fbd75b
	   6 -> #ffb878
	   7 -> #46d6db
	   8 -> #e1e1e1
	   9 -> #5484ed
	   10 -> #51b749
	   11 -> #dc2127
	   Source: Google Colors from Google Calendar API, Colors part
	*/
	activityColorConference  = "10"
	activityColorDefense     = "2"
	activityColorDelivery    = "3"
	activityColorPitch       = "4"
	activityColorProject     = "5"
	activityColorRush        = "6"
	activityColorTP          = "7"
	activityColorKickOff     = "8"
	activityColorProjectTime = "9"
)

var activityColor = map[string]string{
	ActivityConference:  activityColorConference,
	ActivityDefense:     activityColorDefense,
	ActivityDelivery:    activityColorDelivery,
	ActivityPitch:       activityColorPitch,
	ActivityProject:     activityColorProject,
	ActivityRush:        activityColorRush,
	ActivityTP:          activityColorTP,
	ActivityKickOff:     activityColorKickOff,
	ActivityProjectTime: activityColorProjectTime,
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
