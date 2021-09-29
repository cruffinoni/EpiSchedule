package course

import (
	"bytes"
	"fmt"
	"github.com/cruffinoni/EpiSchedule/src/blueprint"
	"github.com/cruffinoni/EpiSchedule/src/environment"
	"io/ioutil"
	"log"
	"net/http"
	"os"
)

func RegisterUserToAnActivity(env environment.Environment, course blueprint.Course, activityId string) {
	for _, activity := range course.Details.Activities {
		if activity.ActivityCode != activityId {
			continue
		}
		if len(activity.Events) == 0 {
			env.Errorf("Activity id %v found but there is no active event. Abort.\n", activity.ActivityCode)
			return
		}
		urlHeader := blueprint.EpitechStartPoint + env.GetAuthentication()
		url := urlHeader + fmt.Sprintf("/module/%v/%v/%v/%v/%v/register?format=json",
			course.Details.Scolaryear, course.Details.Codemodule, course.Details.Codeinstance, activity.ActivityCode, activity.Events[0].Code)
		req, err := http.NewRequest("POST", url, bytes.NewBuffer([]byte("{}")))
		if err != nil {
			log.Fatalf("Err inside new request: '%v'\n", err.Error())
		}
		res, err := env.Client.Do(req)
		if err != nil {
			log.Fatalf("Unable to make http request: '%v'\n", err.Error())
		}
		if res.StatusCode != http.StatusOK {
			env.Errorf("Wanted HTTP code %v but got %v during registering to the event id %v\n",
				http.StatusOK, res.StatusCode, activity.ActivityCode)
			if body, err := ioutil.ReadAll(res.Body); err != nil {
				env.Errorf("Unable to read the body of the request: '%v'\n", err.Error())
			} else {
				env.Errorf("Body: '%v'\n", string(body))
			}
			if res.StatusCode == 500 {
				env.Log(environment.VerboseMedium, "			+ Forbidden. You can't register right now. Try to do it manually\n")
			}
			if env.GetVerboseLevel() != environment.VerboseDebug {
				os.Exit(1)
			}
		} else {
			env.Log(environment.VerboseDebug, "			+ You have been successfully registered to this activity.\n")
		}
	}
}

func RegisterUserToModule(env environment.Environment, module blueprint.Course) {
	urlHeader := blueprint.EpitechStartPoint + env.GetAuthentication()
	url := urlHeader + fmt.Sprintf("/module/%v/%v/%v/register?format=json",
		module.Details.Scolaryear, module.Details.Codemodule, module.Details.Codeinstance)
	req, err := http.NewRequest("POST", url, bytes.NewBuffer([]byte("{}")))
	if err != nil {
		log.Fatalf("Err inside new request: '%v'\n", err.Error())
	}
	res, err := env.Client.Do(req)
	if err != nil {
		log.Fatalf("Unable to make http request: '%v'\n", err.Error())
	}
	if res.StatusCode != http.StatusOK {
		env.Errorf("Wanted HTTP code %v but got %v during registering to the module id %v\n",
			http.StatusOK, res.StatusCode, module.Details.Codemodule)
		if body, err := ioutil.ReadAll(res.Body); err != nil {
			env.Errorf("Unable to read the body of the request: '%v'\n", err.Error())
		} else {
			env.Errorf("Body: '%v'\n", string(body))
		}
		if env.GetVerboseLevel() == environment.VerboseDebug {
			os.Exit(1)
		}
	} else {
		env.Log(environment.VerboseDebug, environment.ColorGreen+"	+ You have been successfully registered to the module.\n")
	}
}
