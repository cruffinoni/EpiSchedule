package planning

import (
	"github.com/Dayrion/EpiSchedule/src/endpoint/course"
	"github.com/Dayrion/EpiSchedule/src/environment"
	"github.com/Dayrion/EpiSchedule/src/utils"
	"log"
)

func ShowIncomingEvents(env environment.Environment) {
	courseList, err := course.GetAllCourses(env)
	if err != nil {
		log.Fatalf("An error occured during retrieving all courses: %v\n", err.Error())
	}
	env.Log(environment.VerboseSimple, "Show all incoming events...\n")
	for _, course := range courseList {
		env.Logf(environment.VerboseSimple, environment.ColorCyan+"~ Show incoming events for module named '%v'\n",
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
				if utils.IsDateAfterNow(activity.End) {
					env.Logf(environment.VerboseSimple, environment.ColorRed+"	! You are registered to the activity [%v] coming from '%v' but the activity is gone.\n",
						activity.Title, course.Summary.Title)
					env.Logf(environment.VerboseSimple, environment.ColorRed+"		! The even ended at %v\n",
						activity.Events[0].End)
					continue
				}
				env.Logf(environment.VerboseSimple, environment.ColorGreen+"	+ You are registered to the activity [%v] coming from '%v'.\n",
					activity.Title, course.Summary.Title)
				if env.IsAutoCalendarRegisteredActivity(activity.TypeTitle) {
					env.Log(environment.VerboseSimple, environment.ColorBlue+"	> This activity will be added to your agenda.\n")
					env.AddEvent(activity)
				}
				env.Logf(environment.VerboseSimple, environment.ColorGreen+"		- The event start at %v and end at %v\n",
					activity.Events[0].Begin, activity.Events[0].End)
			}
		}
	}
}
