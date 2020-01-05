package course

import (
	"blueprint"
	"bytes"
	"endpoint"
	"environment"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

func RegisterUserToAnActivity(env environment.Environment, course blueprint.RegisteredCourse, activityId string) {
	for _, activity := range course.Details.Activites {
		if activity.Codeacti != activityId {
			continue
		}
		if len(activity.Events) == 0 {
			env.Errorf("Activity id %v found but there is no active event. Abort.\n", activity.Codeacti)
			return
		}
		urlHeader := endpoint.EpitechStartPoint + env.GetAuthentication()
		url := urlHeader + fmt.Sprintf("/module/%v/%v/%v/%v/%v/register?format=json",
			course.Details.Scolaryear, course.Details.Codemodule, course.Details.Codeinstance, activity.Codeacti, activity.Events[0].Code)
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
				http.StatusOK, res.StatusCode, activity.Codeacti)
			if body, err := ioutil.ReadAll(res.Body); err != nil {
				env.Errorf("Unable to read the body of the request: '%v'\n", err.Error())
			} else {
				env.Errorf("Body: '%v'\n", string(body))
			}
		} else {
			env.Log(environment.VerboseDebug, "			+ You have been successfully registered to this activity.\n")
		}
	}
}
