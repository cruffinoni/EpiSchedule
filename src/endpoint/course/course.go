package course

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/Dayrion/EpiSchedule/src/blueprint"
	"github.com/Dayrion/EpiSchedule/src/environment"
	"github.com/Dayrion/EpiSchedule/src/utils"
	"io/ioutil"
	"log"
	"net/http"
)

func isAbleToRegister(activity blueprint.CourseActivity) bool {
	return (activity.Register == "1" || activity.RegisterByBloc == "1" ||
		activity.RegisterProf == "1") &&
		(len(activity.Events) > 0 &&
			(activity.Events[0].NbMaxStudentsProjet != "" || activity.Events[0].Location != ""))
}

func checkAllActivitiesFromModule(env environment.Environment, course blueprint.Course) {
	missingOne := false
	for _, activity := range course.Details.Activities {
		if len(activity.Events) == 0 {
			if activity.TypeTitle == "Project" && utils.IsDateBeforeNow(activity.End) {
				env.Logf(environment.VerboseSimple, environment.ColorMagenta+"	!< You are not registered to the main project (named '%v') ! You may register it as soon as possible. The project start at %v.\n",
					activity.Title, activity.Begin)
			} else if utils.IsDateBeforeNow(activity.Begin) && (activity.EndRegister == "" || !isAbleToRegister(activity)) {
				env.Logf(environment.VerboseSimple, environment.ColorBrightYellow+"	!- You are not registered to [%v] but the activity begin to be active at %v. You will be able to register as of this date.\n",
					activity.Title, activity.Begin)
			} else {
				env.Logf(environment.VerboseDebug, environment.ColorBlue+"	~ [%v] might be already passed out or the data is unrecognizable. The activity started at %v\n",
					activity.Title, activity.Begin)
			}
		} else if activity.Events[0].AlreadyRegister == "" {
			//fmt.Printf("Begin: %v & End: %v\n", activity.Begin, activity.End);
			if utils.IsDateAfterNow(activity.End) || utils.IsDateAfterNow(activity.Events[0].End) || (activity.EndRegister == "" && !isAbleToRegister(activity)) {
				env.Logf(environment.VerboseSimple, environment.ColorBrightYellow+"	!- You are not registered to [%v] but the activity is already done or you are unable to register.\n",
					activity.Title)
			} else {
				env.Logf(environment.VerboseSimple, environment.ColorRed+"	!! You are not registered to [%v] typed as [%v] and begin at %v\n",
					activity.Title, activity.TypeTitle, activity.Begin)
				if isAbleToRegister(activity) {
					missingOne = true
					if env.IsAutoRegisteredActivity(activity.TypeTitle) {
						env.Logf(environment.VerboseSimple, environment.ColorCyan+"		~ I'll register you to the activity id %v\n", activity.ActivityCode)
						RegisterUserToAnActivity(env, course, activity.ActivityCode)
					}
				} else {
					env.Logf(environment.VerboseSimple, environment.ColorRed+"		!~ You can't register automatically to this activity because the registrations aren't open: room undefined, no maximum seats specified nor registrations closed.\n"+
						"		!~ You may look the appointments slots on the Epitech's intranet.\n")
				}
			}
		} else {
			env.Logf(environment.VerboseMedium, environment.ColorGreen+"	++ You are registered to [%v] typed as [%v]\n",
				activity.Title, activity.TypeTitle)
		}
	}
	if !missingOne {
		env.Logf(environment.VerboseSimple, environment.ColorCyan+"	+ You are registered for all possible activities in this module.\n")
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
			checkAllActivitiesFromModule(env, course)
		}
	}
}

func getCourseDetails(env environment.Environment, course blueprint.CourseSummary) blueprint.CourseDetails {
	var userCourse blueprint.CourseDetails
	detailsEndpoint := fmt.Sprintf(blueprint.CourseDetailsEndpoint, course.Scolaryear, course.Code, course.Codeinstance)
	if response, err := http.Get(blueprint.EpitechStartPoint + env.GetAuthentication() + detailsEndpoint); err != nil {
		log.Fatal("Invalid response: " + err.Error())
	} else if body, err := ioutil.ReadAll(response.Body); err != nil {
		log.Fatal("Invalid read: " + err.Error())
	} else if err := json.Unmarshal(body, &userCourse); err != nil {
		log.Fatal("Invalid unmarshal: " + err.Error())
	}
	return userCourse
}

func createCoursesList(env environment.Environment, allCourses []blueprint.CourseSummary) []blueprint.Course {
	userCourse := make([]blueprint.Course, 0)
	for _, course := range allCourses {
		if course.Semester < env.User.Semester && (env.User.Semester != 0 && !env.Flag.SpecialSemester) {
			continue
		}
		userCourse = append(userCourse, blueprint.Course{
			Summary: course,
			Details: getCourseDetails(env, course),
		})
	}
	return userCourse
}

func GetAllCourses(env environment.Environment) ([]blueprint.Course, error) {
	var allCourses []blueprint.CourseSummary
	if response, err := http.Get(blueprint.EpitechStartPoint + env.GetAuthentication() + blueprint.CourseDataEndpoint); err != nil || (response != nil && response.StatusCode != http.StatusOK) {
		if err != nil {
			return nil, err
		}
		return nil, errors.New(fmt.Sprintf("got response code %v but wanted %v", response.StatusCode, http.StatusOK))
	} else if body, err := ioutil.ReadAll(response.Body); err != nil {
		return nil, err
	} else if err := json.Unmarshal(body, &allCourses); err != nil {
		return nil, err
	}
	return createCoursesList(env, allCourses), nil
}

func GetCustomCourses(env environment.Environment, endpoint string) ([]blueprint.Course, error) {
	var allCourses []blueprint.CourseSummary
	if response, err := http.Get(endpoint); err != nil || (response != nil && response.StatusCode != http.StatusOK) {
		if err != nil {
			return nil, err
		}
		return nil, errors.New(fmt.Sprintf("got response code %v but wanted %v", response.StatusCode, http.StatusOK))
	} else if body, err := ioutil.ReadAll(response.Body); err != nil {
		return nil, err
	} else if err := json.Unmarshal(body, &allCourses); err != nil {
		return nil, err
	}
	return createCoursesList(env, allCourses), nil
}
