package main

import (
	"endpoint/course"
	"environment"
	"log"
)

func main() {
	env := environment.NewEnvironment(3)
	env.SetVerboseLevel(environment.VerboseDebug)
	env.AddAutoRegisterActivity(environment.ActivityProjectTime, environment.ActivityPitch)
	allCourses, err := course.GetAllCourses(env)
	if err != nil {
		log.Fatalf("An error occured during showing incoming envents: %v\n", err.Error())
	}
	course.ShowNotRegisteredModuleAndActivities(env, allCourses)
}
