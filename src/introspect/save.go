package introspect

import (
	"encoding/json"
	"github.com/Dayrion/EpiSchedule/src/blueprint"
	"github.com/Dayrion/EpiSchedule/src/environment"
	"io/ioutil"
	"log"
	"math/rand"
	"os"
)

const activityPathFile = "./resources/activities.json"

func saveActivities(activities blueprint.Activity) {
	if buffer, err := json.Marshal(activities); err != nil {
		log.Fatalf("Error while marshal: %v\n", err.Error())
	} else {
		file, err := os.OpenFile(activityPathFile, os.O_CREATE|os.O_WRONLY, 0644)
		if err != nil {
			log.Fatalf("Unable to open/create the file 'activities.txt': %v\n", err.Error())
		}
		_, err = file.WriteAt(buffer, 0)
		if err != nil {
			log.Fatalf("An error occurred while writing the activities: %v\n", err.Error())
		}
	}
}

func PopulateActivityType(env environment.Environment, courses []blueprint.Course) {
	env.Log(environment.VerboseDebug, "Retrieving a list of all activities\n")
	file, err := ioutil.ReadFile(activityPathFile)
	var activityList []string
	if err != nil {
		activities := listActivity(courses)
		var actList []string
		for activityName := range activities {
			actList = append(actList, activityName)
		}
		saveActivities(blueprint.Activity{
			Activities: actList,
		})
		activityList = actList
	} else {
		var actList blueprint.Activity
		err := json.Unmarshal(file, &actList)
		if err != nil {
			log.Fatalf("Unable to unmarshal: %v\n", err.Error())
		}
		activityList = actList.Activities
	}
	i := 0
	for _, activity := range activityList {
		switch i {
		case 10:
			environment.AvailableActivity[activity] = "10"
		case 11:
			environment.AvailableActivity[activity] = "11"
			i = 0
		default:
			environment.AvailableActivity[activity] = string(i + 48)
		}
		i++
	}
	env.Logf(environment.VerboseDebug, "%v activity has been found\n", len(activityList))
}

func UpdateActivityList(env environment.Environment, courses []blueprint.Course) {
	env.Log(environment.VerboseDebug, "Updating the activity list...\n")
	update := false
	for _, course := range courses {
		if len(course.Details.Activities) == 0 {
			continue
		}
		activityName := course.Details.Activities[0].TypeTitle
		if _, ok := environment.AvailableActivity[activityName]; !ok {
			environment.AvailableActivity[activityName] = string(rand.Int()%9 + 48)
			update = true
		}
	}
	if update {
		actiList := make([]string, 0)
		for i := range environment.AvailableActivity {
			actiList = append(actiList, i)
		}
		saveActivities(blueprint.Activity{Activities: actiList})
		env.Log(environment.VerboseDebug, "Activity list updated\n")
	} else {
		env.Log(environment.VerboseDebug, "No activity added.\n")
	}
}
