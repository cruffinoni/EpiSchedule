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
	if err != nil {
		log.Fatalf("An error occured during retrieving all courses: %v\n", err.Error())
	}
	flag.SetPreHandlerToCmd(flag.ArgRegister, func() {
		env.AddAutoRegisterActivity(environment.ActivityKickOff, environment.ActivityProjectTime)
	})
	flag.SetPreHandlerToCmd(flag.ArgShow, func() {
		env.SetUpCalendar()
		env.AddAutoRegisterCalendarActivity(environment.ActivityPitch, environment.ActivityKickOff, environment.ActivityProjectTime, environment.ActivityTP)
	})
	flag.SetHandlerToCmd(flag.ArgRegister, course.ShowNotRegisteredModuleAndActivities)
	flag.SetHandlerToCmd(flag.ArgShow, planning.ShowIncomingEvents)
	flag.SetHandlerToCmd(flag.ArgIntrospect, introspect.ListAllActivityFromCourses)

	flag.InitCommandArg(&env)
	cmd := flag.RetrieveCommandFlag(&env, os.Args)
	cmd.ExecuteHandlers(env, allCourses)
}
