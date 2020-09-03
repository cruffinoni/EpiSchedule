package module

import (
	"github.com/Dayrion/EpiSchedule/src/endpoint/course"
	"github.com/Dayrion/EpiSchedule/src/environment"
	"log"
)

func RegisterModuleToCalendar(env environment.Environment) {
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
			env.Log(environment.VerboseSimple, "'%v' has no active project\n", c.Summary.Title)
			continue
		}
		for _, d := range c.Details.Activities {
			if d.TypeTitle == "Project" {
				env.Logf(environment.VerboseSimple, "Project '%v' added. It starts at %v and end at %v\n",
					d.Title, d.Begin, d.End)
				env.AddModule(c.Summary, d)
				break
			}
		}
	}
}
