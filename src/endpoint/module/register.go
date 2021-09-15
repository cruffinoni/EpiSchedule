package module

import (
	"encoding/json"
	"github.com/cruffinoni/EpiSchedule/src/blueprint"
	"github.com/cruffinoni/EpiSchedule/src/endpoint/course"
	"github.com/cruffinoni/EpiSchedule/src/environment"
	"github.com/cruffinoni/EpiSchedule/src/utils"
	"io/ioutil"
	"log"
	"os"
	"time"
)

func getModulesToRegister() []string {
	path, present := os.LookupEnv("MODULE_PATH")
	if !present {
		path = "./resources/module.json"
	}
	file, err := os.OpenFile(path, os.O_RDONLY, 0644)
	if err != nil {
		return []string{}
	}
	all, err := ioutil.ReadAll(file)
	if err != nil {
		return []string{}
	}
	var modules blueprint.ModulesList
	err = json.Unmarshal(all, &modules)
	if err != nil {
		return []string{}
	}
	return modules.Modules
}

func containString(element string, ref []string) bool {
	for _, i := range ref {
		if i == element {
			return true
		}
	}
	return false
}

func RegisterToModule(env environment.Environment) {
	env.Log(environment.VerboseSimple, "Retrieving the list of module that you want to be in.\n")
	modulesToRegister := getModulesToRegister()
	env.Logf(environment.VerboseDebug, "You have registered for %v module(s).\n", len(modulesToRegister))
	env.Log(environment.VerboseSimple, "Getting all modules from the current school year (it may take a while)...\n")
	courseList, err := course.GetAllCourses(env)
	if err != nil {
		log.Fatalf("An error occured during retrieving all courses: %v\n", err.Error())
	}
	env.Log(environment.VerboseSimple, "Trying to register you in the selected module\n")
	for _, c := range courseList {
		if !containString(c.Summary.Title, modulesToRegister) {
			continue
		}
		if c.Details.AllowRegister == 0 {
			dateFormatted, err := utils.GetCalendarDateFromString(c.Details.Begin)
			if err != nil {
				env.Logf(environment.VerboseSimple, environment.ColorBlue+"Module '%v' has not started yet and will begin at '%v'\n", c.Summary.Title, c.Details.Begin)
			} else {
				env.Logf(environment.VerboseSimple, environment.ColorBlue+"Module '%v' has not started yet and will begin at '%v'\n", c.Summary.Title, dateFormatted.Format(time.RFC822))
			}
			continue
		}
		if c.Details.StudentRegistered > 0 {
			env.Logf(environment.VerboseSimple, environment.ColorMagenta+"You are already registered to the module %v\n", c.Summary.Title)
		} else {
			env.Logf(environment.VerboseSimple, environment.ColorCyan+"Module %v found! Registering...\n", c.Summary.Title)
			course.RegisterUserToModule(env, c)
		}
	}
}

func RegisterModuleToCalendar(env environment.Environment) {
	//modulesToRegister := getModulesToRegister()
	env.Log(environment.VerboseSimple, "Getting all courses, it can take some times...\n")
	courseList, err := course.GetAllCourses(env)
	if err != nil {
		log.Fatalf("An error occured during retrieving all courses: %v\n", err.Error())
	}
	env.Log(environment.VerboseSimple, "Registering modules to your calendar\n")
	for _, c := range courseList {
		if c.Summary.Status != "ongoing" {
			dateFormatted, err := utils.GetCalendarDateFromString(c.Details.Begin)
			if err != nil {
				env.Logf(environment.VerboseSimple, environment.ColorBlue+"Module '%v' has not started yet and will begin at '%v'\n", c.Summary.Title, c.Details.Begin)
			} else {
				env.Logf(environment.VerboseSimple, environment.ColorBlue+"Module '%v' has not started yet and will begin at '%v'\n", c.Summary.Title, dateFormatted.Format(time.RFC822))
			}
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
