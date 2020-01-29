package introspect

import (
	"encoding/json"
	"github.com/Dayrion/EpiSchedule/src/blueprint"
	"github.com/Dayrion/EpiSchedule/src/endpoint/course"
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
		_ = os.Remove(activityPathFile)
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

func PopulateActivityType(env environment.Environment, _ []blueprint.Course) {
	env.Log(environment.VerboseDebug, "-> Retrieving a list of all activities\n")
	file, err := ioutil.ReadFile(activityPathFile)
	var activityList []string
	if err != nil {
		env.Log(environment.VerboseDebug, "No activity found. Let's update the resources.\n")
		UpdateActivityList(env, nil)
		env.Log(environment.VerboseDebug, "List created, let's retry...\n")
		PopulateActivityType(env, nil)
		return
	}
	var actList blueprint.Activity
	err = json.Unmarshal(file, &actList)
	if err != nil {
		log.Fatalf("Unable to unmarshal: %v\n", err.Error())
	}
	activityList = actList.Activities
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
	env.Logf(environment.VerboseDebug, "-> %v activity has been found\n", len(activityList))
}

func UpdateActivityList(env environment.Environment, _ []blueprint.Course) {
	env.Log(environment.VerboseDebug, "Updating the activity list, it might take some time...\n")
	update := false
	count := 0
	courses, err := course.GetAllCourses(env)
	if err != nil {
		log.Fatalf("Error while retrieving courses: %v\n", err.Error())
	}
	for _, specificCourse := range courses {
		if len(specificCourse.Details.Activities) == 0 {
			continue
		}
		for _, allActivities := range specificCourse.Details.Activities {
			if _, ok := environment.AvailableActivity[allActivities.TypeTitle]; !ok {
				count++
				environment.AvailableActivity[allActivities.TypeTitle] = string(rand.Int()%9 + 48)
				update = true
				env.Logf(environment.VerboseDebug, "-> Activity %v will be added.\n", allActivities.TypeTitle)
			}
		}
	}
	if update {
		actiList := make([]string, 0)
		for i := range environment.AvailableActivity {
			actiList = append(actiList, i)
		}
		saveActivities(blueprint.Activity{Activities: actiList})
		env.Logf(environment.VerboseDebug, "%v activity added.\n", count)
	} else {
		env.Log(environment.VerboseDebug, "No activity added.\n")
	}
}
