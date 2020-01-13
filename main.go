package main

import (
	//"github.com/Dayrion/EpiSchedule/src/credits"
	"github.com/Dayrion/EpiSchedule/src/endpoint/course"
	"github.com/Dayrion/EpiSchedule/src/endpoint/planning"
	//"github.com/Dayrion/EpiSchedule/src/endpoint/reception"
	"github.com/Dayrion/EpiSchedule/src/environment"
	"github.com/Dayrion/EpiSchedule/src/environment/flag"
	"github.com/Dayrion/EpiSchedule/src/introspect"
	"log"
	"os"
)

func main() {
	env := environment.NewEnvironment()
	env.SetVerboseLevel(environment.VerboseDebug)
	allCourses, err := course.GetAllCourses(env)
	introspect.PopulateActivityType(env, allCourses)
	if err != nil {
		log.Fatalf("An error occured during retrieving all courses: %v\n", err.Error())
	}
	flag.SetUpPreHandler(flag.ArgRegister, func(env *environment.Environment) {
		env.AddAutoRegisterActivity(environment.ActivityKickOff, environment.ActivityProjectTime)
	})
	flag.SetUpPreHandler(flag.ArgShow, func(env *environment.Environment) {
		env.SetUpCalendar()
		env.AddAutoRegisterCalendarActivity(environment.ActivityPitch, environment.ActivityKickOff, environment.ActivityProjectTime, environment.ActivityTP)
	})
	flag.SetHandlerToCmd(flag.ArgRegister, course.ShowNotRegisteredModuleAndActivities)
	flag.SetHandlerToCmd(flag.ArgShow, planning.ShowIncomingEvents)
	flag.SetHandlerToCmd(flag.ArgIntrospect, introspect.ShowActivitiesTypeFromCourses)
	flag.SetHandlerToCmd(flag.ArgUpdate, introspect.UpdateActivityList)
	flag.InitCommandArg(&env)
	cmd := flag.RetrieveCommand(&env, os.Args)
	cmd.ExecuteHandlers(&env, allCourses)
}
