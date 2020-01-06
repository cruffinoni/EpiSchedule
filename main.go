package main

import (
	"github.com/Dayrion/EpiSchedule/src/credits"
	"github.com/Dayrion/EpiSchedule/src/endpoint/course"
	"github.com/Dayrion/EpiSchedule/src/endpoint/planning"
	"github.com/Dayrion/EpiSchedule/src/endpoint/reception"
	"github.com/Dayrion/EpiSchedule/src/environment"
	"github.com/Dayrion/EpiSchedule/src/introspect"
	"log"
	"os"
)

func main() {
	environment.CheckArgs(os.Args)
	env := environment.NewEnvironment()
	env.SetVerboseLevel(environment.VerboseDebug)
	env.RetrieveCommandFlag(os.Args)
	env.User.Semester, env.User.Credits = reception.GetCurrentUserSemesterAndCredits(env)
	credits.DisplayCreditsInfo(env)
	allCourses, err := course.GetAllCourses(env)
	if err != nil {
		log.Fatalf("An error occured during retrieving all courses: %v\n", err.Error())
	}
	if os.Args[1] == environment.FlagRegister {
		env.AddAutoRegisterActivity(environment.ActivityKickOff, environment.ActivityProjectTime)
		course.ShowNotRegisteredModuleAndActivities(env, allCourses)
	} else if os.Args[1] == environment.FlagShow {
		env.SetUpCalendar()
		env.AddAutoRegisterCalendarActivity(environment.ActivityPitch, environment.ActivityKickOff, environment.ActivityProjectTime, environment.ActivityTP)
		planning.ShowIncomingEvents(env, allCourses)
	} else if os.Args[1] == environment.FlagIntrospect {
		introspect.ListAllActivityFromCourses(allCourses)
	}
}
