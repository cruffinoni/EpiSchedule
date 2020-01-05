package planning

import (
	"blueprint"
	"endpoint"
	"environment"
)

func ShowIncomingEvents(env environment.Environment, courseList []blueprint.Course) {
	env.Log(environment.VerboseSimple, "Show all incoming events...\n")
	for _, course := range courseList {
		env.Logf(environment.VerboseSimple, "~ Show incoming events for module named '%v'\n",
			course.Summary.Title)
		if len(course.Details.Activities) == 0 {
			env.Logf(environment.VerboseDebug, "There is no activity for %v\n", course.Details.Title)
			continue
		}
		if course.Details.StudentRegistered == 0 {
			continue
		}
		for _, activity := range course.Details.Activities {
			if len(activity.Events) == 0 {
				continue
			}
			if activity.Events[0].AlreadyRegister != "" {
				if endpoint.IsDateAfterNow(activity.Begin) {
					env.Logf(environment.VerboseSimple, "	! You are registered to the activity [%v] coming from '%v' but the activity is gone.\n",
						activity.Title, course.Summary.Title)
					env.Logf(environment.VerboseSimple, "		! The even ended at %v\n",
						activity.Events[0].End)
					continue
				}
				env.Logf(environment.VerboseSimple, "	+ You are registered to the activity [%v] coming from '%v'.\n",
					activity.Title, course.Summary.Title)
				env.Logf(environment.VerboseSimple, "		- The even start at %v and end at %v\n",
					activity.Events[0].Begin, activity.Events[0].End)
			}
		}
	}
}
