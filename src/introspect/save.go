package introspect

import (
	"encoding/json"
	"fmt"
	"github.com/Dayrion/EpiSchedule/src/blueprint"
	"github.com/Dayrion/EpiSchedule/src/environment"
	"log"
	"os"
)

const activityPathFile = "./resources/activities.json"

func SaveActivities(activities blueprint.Activity) {
	if buffer, err := json.Marshal(activities); err != nil {
		log.Fatalf("Error while marshal: %v\n", err.Error())
	} else {
		fmt.Printf("Buffer ? '%v'\n", string(buffer))
		file, err := os.OpenFile(activityPathFile, os.O_CREATE|os.O_WRONLY, 0644)
		if err != nil {
			log.Fatalf("Unable to open/create the file 'activities.txt': %v\n", err.Error())
		}
		_, err = file.Write(buffer)
		if err != nil {
			log.Fatalf("An error occurred while writing the activities: %v\n", err.Error())
		}
	}
}

func GetAllAvailableActivityType(env environment.Environment, courses []blueprint.Course) {
	file, err := os.Open(activityPathFile)
	if err != nil {
		activities := listActivity(courses)
		var actList []string
		for activityName := range activities {
			actList = append(actList, activityName)
		}
		SaveActivities(blueprint.Activity{
			Activities:actList,
		})
		file, err = os.Open(activityPathFile)
		if err != nil {
			log.Fatalf("Unable to create file for activities: %v\n", err.Error())
		}
	} else {

	}
	file.Name()
}
