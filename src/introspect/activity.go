package introspect

import (
	"fmt"
	"github.com/Dayrion/EpiSchedule/src/blueprint"
	"github.com/Dayrion/EpiSchedule/src/environment"
)

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

func ShowActivitiesTypeFromCourses(env environment.Environment, courses []blueprint.Course) {
	env.Log(environment.VerboseSimple, "List of all activities type:\n")
	activities := listActivity(courses)
	i := 0
	var actList []string
	for activityName := range activities {
		if i == 0 {
			fmt.Printf("%v", activityName)
		} else {
			fmt.Printf(", %v", activityName)
		}
		actList = append(actList, activityName)
		i++
	}
	fmt.Print("\n")
}
