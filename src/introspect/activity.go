package introspect

import (
	"fmt"
	"github.com/Dayrion/EpiSchedule/src/blueprint"
	"github.com/Dayrion/EpiSchedule/src/credits"
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

func ShowActivitiesTypeFromCourses(env environment.Environment) {
	credits.DisplayCreditsInfo(env)
	env.Log(environment.VerboseSimple, "List of all activities type:\n")
	var actList []string
	for activityName := range environment.AvailableActivity {
		fmt.Printf("→ %25v\n", activityName)
		actList = append(actList, activityName)
	}
	fmt.Print("\n")
}
