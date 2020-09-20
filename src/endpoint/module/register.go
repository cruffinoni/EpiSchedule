package module

import (
	"github.com/Dayrion/EpiSchedule/src/endpoint/course"
	"github.com/Dayrion/EpiSchedule/src/environment"
	"log"
)

func RegisterModuleToCalendar(env environment.Environment) {
	env.Log(environment.VerboseSimple, "Getting all courses, it can take some times...\n")
	courseList, err := course.GetAllCourses(env)
	if err != nil {
		log.Fatalf("An error occured during retrieving all courses: %v\n", err.Error())
	}
	env.Log(environment.VerboseSimple, "Registering modules to your calendar\n")
	for _, c := range courseList {
		if c.Summary.Status != "ongoing" {
			continue
		}
		if len(c.Details.Activities) == 0 {
			env.Logf(environment.VerboseSimple, environment.ColorRed+"'%v' has no active project\n", c.Summary.Title)
			continue
		}
		for _, a := range c.Details.Activities {
			if a.TypeTitle == "Project" || a.TypeTitle == "Mini-project" || a.TypeCode == "proj" {
				env.Logf(environment.VerboseSimple, environment.ColorGreen+"Project '%v' added. It starts at %v and end at %v\n",
					a.Title, a.Begin, a.End)
				env.AddModule(c.Summary, a)
				break
			}
		}
	}
}
