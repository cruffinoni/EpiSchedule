package main

import (
	"github.com/cruffinoni/EpiSchedule/src/endpoint/course"
	"github.com/cruffinoni/EpiSchedule/src/endpoint/module"
	"github.com/cruffinoni/EpiSchedule/src/endpoint/planning"
	"github.com/cruffinoni/EpiSchedule/src/endpoint/reception"
	"github.com/cruffinoni/EpiSchedule/src/environment"
	"github.com/cruffinoni/EpiSchedule/src/environment/flag"
	"github.com/cruffinoni/EpiSchedule/src/introspect"
	"os"
)

func setUpCommands(env *environment.Environment) {
	flag.SetHandlerToCmd("register", course.ShowAndRegisterToActivities)
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
		env.Flag.SpecialSemester = true
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
	flag.SetHandlerToCmd("module-calendar", module.RegisterModuleToCalendar)
	flag.SetUpPreHandler("module-calendar", func(e *environment.Environment) {
		e.SetUpCalendar()
	})

	flag.SetHandlerToCmd("module", module.RegisterToModule)
	flag.SetUpPreHandler("module", func(e *environment.Environment) {
		e.SetUpCalendar()
	})
}

func main() {
	env := environment.NewEnvironment()
	env.SetVerboseLevel(environment.VerboseDebug)
	env.User.Semester, env.User.Credits = reception.GetCurrentUserSemesterAndCredits(env)
	introspect.PopulateActivityType(env)
	setUpCommands(&env)
	cmd := flag.RetrieveCommand(&env, os.Args)
	cmd.ExecuteHandlers(&env)
}
