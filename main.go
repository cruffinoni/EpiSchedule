package main

import (
	"fmt"
	"github.com/Dayrion/EpiSchedule/src/credits"
	"github.com/Dayrion/EpiSchedule/src/endpoint/course"
	"github.com/Dayrion/EpiSchedule/src/endpoint/planning"
	"github.com/Dayrion/EpiSchedule/src/endpoint/reception"
	"github.com/Dayrion/EpiSchedule/src/environment"
	"github.com/Dayrion/EpiSchedule/src/environment/flag"
	"github.com/Dayrion/EpiSchedule/src/introspect"
	"log"
	"os"
)

func setUpCommands(env *environment.Environment) {
	flag.SetHandlerToCmd("register", course.ShowNotRegisteredModuleAndActivities)
	flag.SetUpPreHandler("register", func(env *environment.Environment) {
		env.AddAutoRegisterActivity(environment.ActivityToStringArray()...)
	})
	flag.SetArgToCmd("register", flag.ProgArg{
		Hold:         &env.Flag.SpecialSemester,
		DefaultValue: false,
		Name:         "special-semester",
		Description:  "(Optional) Register the semester 0 as a valid one.",
	})

	flag.SetHandlerToCmd("introspect", introspect.ShowActivitiesTypeFromCourses)
	flag.SetArgToCmd("introspect", flag.ProgArg{
		Hold:         &env.Flag.SpecialSemester,
		DefaultValue: true,
		Name:         "special-semester",
		Description:  "(Optional) Register the semester 0 as a valid one. It will give more type.",
	})

	flag.SetHandlerToCmd("show", planning.ShowIncomingEvents)
	flag.SetUpPreHandler("show", func(env *environment.Environment) {
		fmt.Printf("Going right to prehanlder show\n");
		env.SetUpCalendar()
		env.AddAutoRegisterCalendarActivity(environment.ActivityToStringArray()...)
	})
	flag.SetHandlerToCmd("update", introspect.UpdateActivityList)
	flag.SetArgToCmd("update", flag.ProgArg{
		Hold:         &env.Flag.SpecialSemester,
		DefaultValue: true,
		Name:         "special-semester",
		Description:  "(Optional) Register the semester 0 as a valid one. It will give more type.",
	})
}

func main() {
	env := environment.NewEnvironment()
	env.SetVerboseLevel(environment.VerboseDebug)
	env.User.Semester, env.User.Credits = reception.GetCurrentUserSemesterAndCredits(env)
	allCourses, err := course.GetAllCourses(env)
	if err != nil {
		log.Fatalf("An error occured during retrieving all courses: %v\n", err.Error())
	}
	credits.DisplayCreditsInfo(env)
	introspect.PopulateActivityType(env, allCourses)
	setUpCommands(&env)
	cmd := flag.RetrieveCommand(&env, os.Args)
	cmd.ExecuteHandlers(&env, allCourses)
}
