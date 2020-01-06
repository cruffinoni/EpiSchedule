package introspect

import (
	"fmt"
	"github.com/Dayrion/EpiSchedule/src/blueprint"
)

func ListAllActivityFromCourses(courses []blueprint.Course) {
	fmt.Print("List of all activities type:\n")
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
	for i := range activities {
		fmt.Printf("%v\n", i)
	}
}
