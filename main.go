package main

import (
	"endpoint/course"
	"endpoint/planning"
	"environment"
	"log"
	"os"
)

func main() {
	env := environment.NewEnvironment(3)
	env.AddAutoRegisterActivity(environment.ActivityPitch)
	env.SetVerboseLevel(environment.VerboseDebug)
	if len(os.Args) != 2 {
		log.Fatalf("Invalid args count. Wanted 1 but got %v instead.\n", len(os.Args)-1)
	}
	allCourses, err := course.GetAllCourses(env)
	if err != nil {
		log.Fatalf("An error occured during retrieving all courses: %v\n", err.Error())
	}
	if os.Args[1] == "register" {
		env.AddAutoRegisterActivity(environment.ActivityPitch)
		course.ShowNotRegisteredModuleAndActivities(env, allCourses)
	} else if os.Args[1] == "show" {
		planning.ShowIncomingEvents(env, allCourses)
	} else {
		log.Fatalf("'%v' is an invalid arg\n", os.Args[1])
	}
}
