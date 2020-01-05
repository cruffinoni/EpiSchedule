package course

import (
	"blueprint"
	"encoding/json"
	"endpoint"
	"environment"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

func isAbleToRegister(activity blueprint.CourseActivity) bool {
	return activity.Register == "1" || activity.RegisterByBloc == "1" ||
		activity.RegisterProf == "1"
}

func checkAllActivitiesFromModule(env environment.Environment, course blueprint.Course) {
	missingOne := false
	for _, activity := range course.Details.Activities {
		if len(activity.Events) == 0 {
			if activity.TypeTitle == "Project" && endpoint.IsDateBeforeNow(activity.End) {
				env.Logf(environment.VerboseSimple, "	!< You are not registered to the main project (named '%v') ! You may register it as soon as possible. The project start at %v.\n",
					activity.Title, activity.Begin)
			} else if endpoint.IsDateBeforeNow(activity.Begin) && (activity.EndRegister == "" || !isAbleToRegister(activity)) {
				env.Logf(environment.VerboseSimple, "	!- You are not registered to [%v] but the activity begin to be active at %v. You will be able to register as of this date.\n",
					activity.Title, activity.Begin)
			} else {
				env.Logf(environment.VerboseDebug, "	~ [%v] might be already passed out or the data is unrecognizable. The activity started at %v\n",
					activity.Title, activity.Begin)
			}
		} else if activity.Events[0].AlreadyRegister == "" {
			missingOne = true
			env.Logf(environment.VerboseSimple, "	!! You are not registered to [%v] which is: [%v]\n",
				activity.Title, activity.TypeTitle)
			if env.IsAutoRegisteredActivity(activity.TypeTitle) {
				if !isAbleToRegister(activity) {
					env.Logf(environment.VerboseSimple, "		!~ You can't register automatically to this activity because the registrations aren't open.\n		!~ You may look the appointments slots on the Epitech's intranet.\n")
				} else {
					env.Logf(environment.VerboseSimple, "		~ I'll register you to the activity id %v\n", activity.ActivityCode)
					RegisterUserToAnActivity(env, course, activity.ActivityCode)
				}
			}
		} else {
			env.Logf(environment.VerboseMedium, "	++ You are registered to [%v] typed as [%v]\n",
				activity.Title, activity.TypeTitle)
		}
	}
	if !missingOne {
		env.Logf(environment.VerboseSimple, "	+ You are registered to all activities for this module.\n")
	}
}

func ShowNotRegisteredModuleAndActivities(env environment.Environment, courses []blueprint.Course) {
	if env.GetVerboseLevel() < environment.VerboseSimple {
		return
	}
	for _, course := range courses {
		if course.Details.Opened != "1" {
			env.Logf(environment.VerboseMedium, "The course '%v' is in your semester but it's closed\n", course.Details.Title)
		}
		if course.Details.StudentRegistered == 0 {
			env.Logf(environment.VerboseSimple, "! You seems not to be registered to the module: %v\n", course.Details.Title)
		} else {
			env.Logf(environment.VerboseSimple, "+ You are registered to the module: %v\n", course.Details.Title)
		}
		checkAllActivitiesFromModule(env, course)
	}
}

func getCourseDetails(env environment.Environment, course blueprint.CourseSummary) blueprint.CourseDetails {
	var userCourse blueprint.CourseDetails
	detailsEndpoint := fmt.Sprintf(blueprint.CourseDetailsEndpoint, course.Scolaryear, course.Code, course.Codeinstance)
	if response, err := http.Get(endpoint.EpitechStartPoint + env.GetAuthentication() + detailsEndpoint); err != nil {
		log.Fatal("Invalid response: " + err.Error())
	} else if body, err := ioutil.ReadAll(response.Body); err != nil {
		log.Fatal("Invalid read: " + err.Error())
	} else if err := json.Unmarshal(body, &userCourse); err != nil {
		log.Fatal("Invalid unmarshal: " + err.Error())
	}
	return userCourse
}

func GetAllCourses(env environment.Environment) ([]blueprint.Course, error) {
	var allCourses []blueprint.CourseSummary
	if response, err := http.Get(endpoint.EpitechStartPoint + env.GetAuthentication() + blueprint.CourseDataEndpoint); err != nil {
		return nil, err
	} else if body, err := ioutil.ReadAll(response.Body); err != nil {
		return nil, err
	} else if err := json.Unmarshal(body, &allCourses); err != nil {
		return nil, err
	}
	userCourse := make([]blueprint.Course, 0)
	for _, course := range allCourses {
		if course.Semester < env.User.Semester {
			continue
		}
		userCourse = append(userCourse, blueprint.Course{
			Summary: course,
			Details: getCourseDetails(env, course),
		})
	}
	return userCourse, nil
}
