package introspect

import (
	"fmt"
	"github.com/Dayrion/EpiSchedule/src/blueprint"
	"github.com/Dayrion/EpiSchedule/src/environment"
	"log"
	"os"
)

const activityPathFile = "./resources/activities.txt"

func SaveActivities(activities string) {
	file, err := os.OpenFile(activityPathFile, os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		log.Fatalf("Unable to open/create the file 'activities.txt': %v\n", err.Error())
	}
	_, err = file.WriteString(activities)
	if err != nil {
		log.Fatalf("An error occurred while writing the activities: %v\n", err.Error())
	}
}

func listActivity(courses []blueprint.Course) map[string]int {
	activities := make(map[string]int)
	for _, course := range courses {
		if len(course.Details.Activities) == 0 {
			continue
		}
		activityName := course.Details.Activities[0].TypeTitle
		if activities[activityName] == 0 {
			activities[activityName]++
		}
	}
	return activities
}

func GetAllAvailableActivityType(env environment.Environment, courses []blueprint.Course) {
	file, err := os.Open(activityPathFile)
	if err != nil {
		activities := listActivity(courses)
		actList := ""
		for activityName := range activities {
			actList += activityName + "\n"
		}
		SaveActivities(actList)
		file, err = os.Open(activityPathFile)
		if err != nil {
			log.Fatalf("Unable to create file for activities: %v\n", err.Error())
		}
	}
	file.Name()
}

func ListAllActivityFromCourses(env environment.Environment, courses []blueprint.Course) {
	env.Log(environment.VerboseSimple, "List of all activities type:\n")
	activities := listActivity(courses)
	i := 0
	actList := ""
	for activityName := range activities {
		if i == 0 {
			fmt.Printf("%v", activityName)
		} else {
			fmt.Printf(", %v", activityName)
		}
		actList += activityName + "\n"
		i++
	}
	if env.Flag.SaveActivities {
		env.Log(environment.VerboseSimple, "Registering all activities.\n")
		SaveActivities(actList)
		env.Log(environment.VerboseSimple, "Activities saved.\n")
	}
}
