package module

import (
	"encoding/json"
	"endpoint"
	"environment"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"time"
)

func ShowNotRegisteredActivities(modules []ModuleStruct, moduleCode string) {
	missingOne := false
	for _, module := range modules {
		if moduleCode == module.ModuleCode {
			if !module.isUserRegistered() {
				missingOne = true
				startActive, err := endpoint.GetDateFromString(module.BeginActive)
				//fmt.Printf("-> %+v\n", module)
				if err != nil {
					log.Fatalf("Unable to retrieve date from string inside (ShowNotRegisteredActivities): %v\n", err.Error())
				}
				if module.TypeActivity == "Project" {
					fmt.Printf("	>! You are not registered to the project [%v] ! You may register it as soon as possible.\n",
						module.ActivityTitle)
				} else if time.Now().Before(startActive) && module.BeginEvent == "" {
					fmt.Printf("	! You are not registered to [%v] but the activity begin to be active at %v\n",
						module.ActivityTitle, module.BeginActive)
				} else {
					fmt.Printf("	!! You are not registered to [%v] which is: [%v]\n",
						module.ActivityTitle, module.TypeActivity)
				}
			}
		}
	}
	if !missingOne {
		fmt.Print("	+ You are registered to all activities for this module.\n")
	}
}

func ShowIncomingEvent(modules []ModuleStruct) error {
	fmt.Print("- List of incoming event(s):\n")
	for _, module := range modules {
		if !module.isModuleStarted() {
			if module.BeginEvent != "" {
				fmt.Printf("~ Event [%v] is incoming in [%v]\n", module.ActivityTitle, module.timeLeftBeforeModuleStart()) //start.Day() - time.Now().Day())
			} else {
				fmt.Printf("~ Event [%v] is incoming. The module has no available starting date.\n", module.ActivityTitle)
			}
			fmt.Printf("  ~ This event is [%v]\n", module.TypeActivity)
		}
	}
	return nil
}

func NewModulesBoard(env environment.Environment, start, end time.Time) ([]ModuleStruct, error) {
	var userModule []ModuleStruct
	begging := fmt.Sprintf("%d-%02d-%02d", start.Year(), start.Month(), start.Day())
	ending := fmt.Sprintf("%d-%02d-%02d", end.Year(), end.Month(), end.Day())
	fmt.Printf("Creating a modules board from %v to %v\n", begging, ending)
	if response, err := http.Get(endpoint.EpitechStartPoint + env.GetAuthentication() + fmt.Sprintf(ModuleEndPoint, begging, ending)); err != nil {
		return nil, err
	} else if body, err := ioutil.ReadAll(response.Body); err != nil {
		return nil, err
	} else if err := json.Unmarshal(body, &userModule); err != nil {
		return nil, err
	}
	return userModule, nil
}

/*

	endpoint := EpitechStartPoint + user.GetAuthentication() + fmt.Sprintf("/module/2019/%v/%v/acti-376004/event-368655/unregister?format=json", module.ModuleCode, module.Codeinstance)
	req, err := http.NewRequest("POST", endpoint, bytes.NewBuffer([]byte("{}")))
	if err != nil {
		log.Fatalf("Err inside new request: '%v'\n", err.Error())
	}
	res, err := http.DefaultClient.Do(req)

	if err != nil {
		log.Fatalf("Error: '%v'\n", err.Error())
	}

	fmt.Println("response Status:", res.StatusCode)
	body, _ := ioutil.ReadAll(res.Body)
	fmt.Println("response Body:", string(body))
*/
